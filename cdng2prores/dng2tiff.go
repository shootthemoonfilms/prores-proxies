package main

import (
	pnm "github.com/jbuchbinder/gopnm"
	tiff "golang.org/x/image/tiff"
	"os"
	"os/exec"
)

// DngToTiff handles the conversion process using dcraw and pnmtotiff
func DngToTiff(in, out string) error {
	w, err := os.Create(out)
	if err != nil {
		return err
	}
	defer w.Close()

	c1 := exec.Command(*dcrawPath, "-c", "-q", "0", in)
	c2 := exec.Command(*pnmtotiffPath, "-quiet", "-lzw")
	c2.Stdin, err = c1.StdoutPipe()
	if err != nil {
		return nil
	}
	c2.Stdout = w

	c2.Start()
	c1.Run()
	c2.Wait()

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
	c1.Stdout = w

	// Grab dcraw output
	p, err := c1.StdoutPipe()
	if err != nil {
		return err
	}

	c1.Start()
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
