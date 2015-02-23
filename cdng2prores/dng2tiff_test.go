package main

import (
	// "os"
	"testing"
)

const (
	testFrame = "testframes/Blackmagic Pocket Cinema Camera_1_2013-11-15_1740_C0000_000545.dng"
)

func BenchmarkNonNative(b *testing.B) {
	err := DngToTiff(testFrame, "non-native.tiff")
	//defer os.Remove("non-native.tiff")
	if err != nil {
		b.Log(err)
		b.Fail()
	}
}

func BenchmarkNative(b *testing.B) {
	err := DngToTiffNative(testFrame, "native.tiff")
	//defer os.Remove("native.tiff")
	if err != nil {
		b.Log(err)
		b.Fail()
	}
}

