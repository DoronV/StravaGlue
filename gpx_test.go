package main

import (
	"strings"
	"testing"
)

func TestOpeningGpxFile(t *testing.T) {
	gpxFilePath := "test_files/2021-06-02_382782068_Fietsen.gpx"

	gpxString := Read(gpxFilePath)

	t.Run("can open a gpx file from a path", func(t *testing.T) {
		startString := "<?xml version='1.0' encoding='UTF-8'?>"

		if !strings.Contains(gpxString, startString) {
			t.Errorf("got %s, want %s", gpxString, startString)
		}
	})

	t.Run("reads to the end of the gpx file", func(t *testing.T) {
		endString := "</gpx>"

		if !strings.Contains(gpxString, endString) {
			t.Errorf("got %s, want %s", gpxString, endString)
		}
	})
}

func TestPrepareFiles(t *testing.T) {

	t.Run("can prepare primary file", func(t *testing.T) {
		gpxFilePath := "test_files/Early.gpx"

		gpxString := Read(gpxFilePath)
		got := PreparePrimaryFile(gpxString)

		notWant := "    </trkseg>\n  </trk>\n</gpx>"

		if strings.Contains(got, notWant) {
			t.Errorf("got %s, don't want %s", got, notWant)
		}
	})

	t.Run("can prepare secondary file", func(t *testing.T) {
		gpxFilePath := "test_files/Later.gpx"

		gpxString := Read(gpxFilePath)
		got := PrepareSecondaryFile(gpxString)
		notWant := "</name>"

		if strings.Contains(got, notWant) {
			t.Errorf("got %s, don't want %s", got, notWant)
		}
	})
}
