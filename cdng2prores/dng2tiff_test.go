package main

import (
	// "os"
	"testing"
)

const (
	testFrame = "testframes/000545.dng"
)

func BenchmarkNonNative(b *testing.B) {
	*dcrawPath = "/usr/bin/dcraw" // HACK! HACK!
	err := DngToTiff(testFrame, "frame.tiff")
	//defer os.Remove("non-native.tiff")
	if err != nil {
		b.Log(err)
		b.Fail()
	}
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
