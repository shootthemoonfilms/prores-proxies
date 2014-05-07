package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func convert(pathName, fileName, tmpPath, ffmpegPath string, scaleW, scaleH int) (string, error) {
	var scalingParameter string
	if scaleW > 0 && scaleH > 0 {
		scalingParameter = fmt.Sprintf("-filter:v scale=%d:%d", scaleW, scaleH)
	} else {
		scalingParameter = ""
	}

	log.Print("Processing " + fileName + " in '" + pathName + "'")
	outPath := tmpPath + string(os.PathSeparator) + fileName
	_ = os.MkdirAll(tmpPath, 0755)
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
		Path: ffmpegPath,
		Args: args,
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Start(); err != nil {
		log.Print(err.Error())
		return outPath, err
	}
	if err := command.Wait(); err != nil {
		log.Print(err.Error())
		return outPath, err
	}
	return outPath, nil
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
