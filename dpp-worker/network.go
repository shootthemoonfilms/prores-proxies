package main

import (
	"net/http"
)

func GrabFile(urlBase string, originalFile string, outFile string) error {
	req, err := http.NewRequest(urlBase+"/"+originalFile, url, nil)
	if err != nil {
		return err
	}
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(outFile)
	defer out.Close()

	_, err := io.Copy(out, res.Body)

	return err
}
