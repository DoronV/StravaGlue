package main

import (
	"strings"
	"testing"
)

func TestOpeningGpxFile(t *testing.T) {
	gpxFilePath := "2021-06-02_382782068_Fietsen.gpx"

	gpxReader := Read(gpxFilePath)

	t.Run("can open a gpx file from a path", func(t *testing.T) {
		startString := "<?xml version='1.0' encoding='UTF-8'?>"

		if !strings.Contains(string(gpxReader.GpxString), startString) {
			t.Errorf("got %s, want %s", gpxReader, startString)
		}
	})

	t.Run("reads to the end of the gpx file", func(t *testing.T) {
		endString := "</gpx"

		if !strings.Contains(string(gpxReader.GpxString), endString) {
			t.Errorf("got %s, want %s", gpxReader, endString)
		}
	})
}