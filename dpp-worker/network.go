package main

import (
	"io"
	"net/http"
	"os"
)

func GrabFile(urlBase string, originalFile string, outFile string) error {
	res, err := http.Get(urlBase + "/" + originalFile)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(outFile)
	defer out.Close()

	_, err = io.Copy(out, res.Body)

	return err
}
