package main

import (
	pnm "github.com/jbuchbinder/gopnm"
	tiff "golang.org/x/image/tiff"
	//"log"
	"os"
	"os/exec"
)

// DngToTiff handles the conversion process using dcraw
func DngToTiff(in, out string) error {
	w, err := os.Create(out)
	if err != nil {
		return err
	}
	defer w.Close()

	c1 := exec.Command(*dcrawPath, "-c", "-q", "0", "-T", in)
	c1.Stdout = w
	err = c1.Start()
	if err != nil {
		return err
	}
	c1.Wait()

	return nil
}

// DngToTiffNative handles the conversion process using dcraw
func DngToTiffNative(in, out string) error {
	// File writer
	w, err := os.Create(out)
	if err != nil {
		return err
	}
	defer w.Close()

	c1 := exec.Command(*dcrawPath, "-c", "-q", "0", in)

	// Grab dcraw output
	p, err := c1.StdoutPipe()
	if err != nil {
		return err
	}

	err = c1.Start()
	if err != nil {
		return err
	}
	c1.Wait()

	img, err := pnm.Decode(p)
	if err != nil {
		return err
	}
	err = tiff.Encode(w, img, &tiff.Options{
		Compression: tiff.Deflate,
		Predictor:   true,
	})
	if err != nil {
		return err
	}
	return nil
}
