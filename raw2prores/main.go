package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

var (
	dcrawPath        = flag.String("dcraw", "."+string(os.PathSeparator)+"dcraw", "Path to DCRAW binary")
	ffmpegPath       = flag.String("ffmpeg", "."+string(os.PathSeparator)+"ffmpeg", "Path to FFMPEG binary")
	pnmtotiffPath    = flag.String("pnmtotiff", "."+string(os.PathSeparator)+"pnmtotiff", "Path to PNMTOTIFF binary")
	proxyDir         = flag.String("proxy", "proxy", "Proxy files subdirectory name")
	scaleW           = flag.Int("scalew", 0, "Scale width")
	scaleH           = flag.Int("scaleh", 0, "Scale height")
	extension        = flag.String("extension", "mov", "File extension")
	formatExtension  = flag.String("format-extension", "mov", "File extension used for determining format")
	scalingParameter string
	wg               sync.WaitGroup
)

func main() {
	flag.Parse()

	args := flag.Args()

	if *scaleW > 0 && *scaleH > 0 {
		scalingParameter = fmt.Sprintf("-filter:v scale=%d:%d", *scaleW, *scaleH)
	} else {
		scalingParameter = ""
	}

	log.Printf("Setting maximum parallelism at %d threads", MaxParallelism())
	runtime.GOMAXPROCS(MaxParallelism())

	// Determine if we're using local directory or list of provided ones.
	var dirs []string
	if len(args) < 1 {
		log.Print("Using current working directory to scan")
		dirs = make([]string, 1)
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dirs[0] = cwd
		log.Print("Found cwd " + cwd)
	} else {
		dirs = args[:]
	}

	// Iterate through directories
	for idx := range dirs {
		scanDir(dirs[idx])
	}

	log.Print("Waiting for threads to finish")
	wg.Wait()
	log.Print("Run completed")
}

func scanDir(dirName string) {
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		//log.Print(f.Name())
		fullPath := dirName + string(os.PathSeparator) + f.Name()
		if FileExists(fullPath) && strings.HasSuffix(f.Name(), "."+*formatExtension) {
			log.Print("Processing " + f.Name())
			wg.Add(1)
			go processFile(dirName, f.Name())
		}
	}
}

func processFile(pathName, fileName string) {
	defer wg.Done()

	log.Print("Processing " + fileName + " in '" + pathName + "'")

	// Spin up threads to convert DNG -> TIF by way of pnmtotiff

	outPath := pathName + string(os.PathSeparator) + *proxyDir + string(os.PathSeparator) + fileName
	_ = os.MkdirAll(pathName+string(os.PathSeparator)+*proxyDir, 0755)
	args := []string{
		"-y",
		"-i", pathName + string(os.PathSeparator) + fileName,
	}
	if scalingParameter != "" {
		args = append(args, scalingParameter)
	}

	startNumber := 0 // TODO : calculate from file names

	args = append(args, "-start_number", fmt.Sprintf("%d", startNumber))

	// Specify frame rate
	args = append(args, "-r", "23.97") // TODO: allow specifying this

	// Ultrafast x264+AAC
	args = append(args, "-vcodec", "prores")
	args = append(args, "-profile", "3")
	//args = append(args, "-acodec", "aac")

	args = append(args, outPath)
	command := exec.Cmd{
		Path: *ffmpegPath,
		Args: args,
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Start(); err != nil {
		log.Print(err.Error())
		return
	}
	if err := command.Wait(); err != nil {
		log.Print(err.Error())
		return
	}

	// If everything works, rename if necessary
	if *formatExtension != *extension {
		modifiedFileName := strings.TrimRight(fileName, "."+*formatExtension) + "." + *extension
		os.Rename(outPath, pathName+string(os.PathSeparator)+*proxyDir+string(os.PathSeparator)+modifiedFileName)
	}
}

// DngToTiff handles the conversion process using dcraw and pnmtotiff
func DngToTiff(in, out string) {
	w, _ := os.Create(out)
	defer w.Close()

	c1 := exec.Command(*dcrawPath, "-c", "-q", "0", in)
	c2 := exec.Command(*pnmtotiffPath, "-quiet", "-lzw")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = w

	c2.Start()
	c1.Run()
	c2.Wait()
}

// FileExists reports whether the named file exists.
func FileExists(name string) bool {
	st, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if st.IsDir() {
		return false
	}
	return true
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
