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
	ffmpegPath       = flag.String("ffmpeg", "."+string(os.PathSeparator)+"ffmpeg", "Path to FFMPEG binary")
	proxyDir         = flag.String("proxy", "proxy", "Proxy files subdirectory name")
	scaleW           = flag.Int("scalew", 0, "Scale width")
	scaleH           = flag.Int("scaleh", 0, "Scale height")
	extension        = flag.String("extension", "mov", "File extension")
	threading        = flag.Bool("threading", false, "Use multi-threading")
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

	if *threading {
		log.Print("Setting maximum parallelism")
		runtime.GOMAXPROCS(MaxParallelism())
	}

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

	if *threading {
		log.Print("Waiting for threads to finish")
		wg.Wait()
		log.Print("Run completed")
	}
}

func scanDir(dirName string) {
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		//log.Print(f.Name())
		fullPath := dirName + string(os.PathSeparator) + f.Name()
		if FileExists(fullPath) && strings.HasSuffix(f.Name(), ".mov") {
			log.Print("Processing " + f.Name())
			if *threading {
				wg.Add(1)
				go processFile(dirName, f.Name())
			} else {
				processFile(dirName, f.Name())
			}
		}
	}
}

func processFile(pathName, fileName string) {
	if *threading {
		defer wg.Done()
	}
	log.Print("Processing " + fileName + " in '" + pathName + "'")
	outPath := pathName + string(os.PathSeparator) + *proxyDir + string(os.PathSeparator) + fileName
	_ = os.MkdirAll(pathName+string(os.PathSeparator)+*proxyDir, 0755)
	args := []string{
		"-y",
		"-i", pathName + string(os.PathSeparator) + fileName,
	}
	if scalingParameter != "" {
		args = append(args, scalingParameter)
	}

	// Ultrafast x264+AAC
	args = append(args, "-vcodec", "libx264")
	args = append(args, "-acodec", "aac")
	args = append(args, "-strict", "-2")
	args = append(args, "-preset", "ultrafast")

	// Baseline compatibility, all devices
	args = append(args, "-profile:v", "baseline", "-level", "3.0")

	// YUV 4:2:0
	args = append(args, "-pix_fmt", "yuv420p")

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
	modifiedFileName := strings.TrimRight(fileName, ".mov") + "." + *extension
	os.Rename(outPath, pathName+string(os.PathSeparator)+*proxyDir+string(os.PathSeparator)+modifiedFileName)
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
