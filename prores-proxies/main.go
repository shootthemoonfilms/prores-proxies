package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	ffmpegPath = flag.String("ffmpeg", "."+string(os.PathSeparator)+"ffmpeg", "Path to FFMPEG binary")
	proxyDir   = flag.String("proxy", "proxy", "Proxy files subdirectory name")
)

func main() {
	flag.Parse()

	args := flag.Args()

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
}

func scanDir(dirName string) {
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		//log.Print(f.Name())
		fullPath := dirName + string(os.PathSeparator) + f.Name()
		if FileExists(fullPath) && strings.HasSuffix(f.Name(), ".mov") {
			log.Print("Processing " + f.Name())
			processFile(dirName, f.Name())
		}
	}
}

func processFile(pathName, fileName string) {
	log.Print("Processing " + fileName + " in '" + pathName + "'")
	outPath := pathName + string(os.PathSeparator) + *proxyDir + string(os.PathSeparator) + fileName
	_ = os.MkdirAll(pathName+string(os.PathSeparator)+*proxyDir, 0755)
	command := exec.Cmd{
		Path: *ffmpegPath,
		Args: []string{
			"-y",
			"-i", pathName + string(os.PathSeparator) + fileName,
			"-vcodec", "libx264",
			"-acodec", "aac",
			"-strict", "-2",
			"-preset", "ultrafast",
			outPath,
		},
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
