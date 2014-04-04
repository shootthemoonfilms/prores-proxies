package main

import (
	"flag"
	"github.com/mikespook/gearman-go/worker"
	"github.com/mikespook/golib/signal"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var (
	ffmpeg = flag.String("ffmpeg", "/usr/bin/ffmpeg", "Path to ffmpeg binary")
	temp   = flag.String("temp", "/tmp", "Temporary directory")
)

func CreateProxy(job worker.Job) ([]byte, error) {
	log.Printf("CreateProxy: Data=[%s]\n", job.Data())

	// Pull job definition from JSON representation
	var j ProxyJob
	err := json.Unmarshal(job.Data(), &j)
	if err != nil {
		return nil, err
	}

	// Pull file from the specified URL to the temporary directory
	fn := *temp + "/" + j.OriginalFile + ".proxy.mp4"
	log.Printf("CreateProxy: Grabbing " + j.UrlBase + "/" + j.OriginalFile + " into " + fn)
	err = GrabFile(j.UrlBase, j.OriginalFile, fn)
	log.Printf("CreateProxy: Finished retrieving file")

	for i := 0; i < 10; i++ {
		job.SendWarning([]byte{byte(i)})
		job.SendData([]byte{byte(i)})
		job.UpdateStatus(i+1, 100)
	}
	return job.Data(), nil
}

func main() {
	flag.Parse()

	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")
	w := worker.New(worker.Unlimited)
	defer w.Close()
	w.ErrorHandler = func(e error) {
		log.Println(e)
		if opErr, ok := e.(*net.OpError); ok {
			if !opErr.Temporary() {
				proc, err := os.FindProcess(os.Getpid())
				if err != nil {
					log.Println(err)
				}
				if err := proc.Signal(os.Interrupt); err != nil {
					log.Println(err)
				}
			}
		}
	}
	w.JobHandler = func(job worker.Job) error {
		log.Printf("Data=%s\n", job.Data())
		return nil
	}
	w.AddServer("tcp4", "127.0.0.1:4730")
	w.AddFunc("CreateProxy", CreateProxy, worker.Unlimited)
	if err := w.Ready(); err != nil {
		log.Fatal(err)
		return
	}
	go w.Work()
	sh := signal.NewHandler()
	sh.Bind(os.Interrupt, func() bool { return true })
	sh.Loop()
}
