package main

import (
	"os"
	"sync"
	"testing"
)

var (
	testFrames = []string{
		"testframes/000545.dng",
		"testframes/000546.dng",
		"testframes/000547.dng",
		"testframes/000548.dng",
	}
)

func BenchmarkNonNative(b *testing.B) {
	*dcrawPath = "/usr/bin/dcraw" // HACK! HACK!

	var wg sync.WaitGroup

	b.Logf("Spinning up %d frames", len(testFrames))
	for i := 0; i < len(testFrames); i++ {
		wg.Add(1)
		go func(testFrame string) {
			b.Log("Spinning up thread to process " + testFrame)
			defer wg.Done()
			err := DngToTiff(testFrame, testFrame+".tiff")
			defer os.Remove(testFrame + ".tiff")
			if err != nil {
				b.Log(err)
				//b.Fail()
			}
		}(testFrames[i])
	}

	b.Log("Waiting for threads to finish executing")
	wg.Wait()
}

/*
func BenchmarkNative(b *testing.B) {
	*dcrawPath = "/usr/bin/dcraw"
	err := DngToTiffNative(testFrame, "native.tiff")
	//defer os.Remove("native.tiff")
	if err != nil {
		b.Log(err)
		b.Fail()
	}
}
*/
