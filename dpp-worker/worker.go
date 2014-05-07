package main

import (
	"encoding/json"
	"flag"
	"github.com/mikespook/gearman-go/worker"
	"github.com/mikespook/golib/signal"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	//"strings"
	//"time"
)

var (
	ffmpeg         = flag.String("ffmpeg", "/usr/bin/ffmpeg", "Path to ffmpeg binary")
	temp           = flag.String("temp", "/tmp", "Temporary directory")
	localIPAddress string
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

	log.Printf("CreateProxy: Beginning conversion process")
	outPath, err := convert(*temp, j.OriginalFile, *temp+string(os.PathSeparator)+"proxy", *ffmpeg, j.ResolutionW, j.ResolutionH)
	var errText string
	if err != nil {
		errText = err.Error()
	}

	var filesize int64
	stat, err := os.Stat(outPath)
	if err != nil {
		if errText != "" {
			errText += "\n"
		}
		errText += err.Error()
	} else {
		filesize = stat.Size()
	}

	res := ProxyJobResult{
		Url:   "http://" + localIPAddress + ":4731/proxy/" + path.Base(outPath),
		Size:  filesize,
		Error: errText,
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		job.SendData(b)
		job.UpdateStatus(0, 100)
	}

	return job.Data(), nil
}

func main() {
	flag.Parse()

	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")

	var err error
	localIPAddress, err = hostIP()
	if err != nil {
		panic(err)
	}

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
	w.AddServer("tcp4", ":4730")
	w.AddFunc("CreateProxy", CreateProxy, worker.Unlimited)
	if err := w.Ready(); err != nil {
		log.Fatal(err)
		return
	}
	go w.Work()

	// Add result file server
	go http.ListenAndServe(":4731", http.FileServer(http.Dir(*temp)))

	sh := signal.NewHandler()
	sh.Bind(os.Interrupt, func() bool { return true })
	sh.Loop()
}
