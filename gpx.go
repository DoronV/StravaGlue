package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Read(filepath string) string {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Errorf("error reading the file, %v", err)
	}

	GpxString := string(data)

	return GpxString
}

func PreparePrimaryFile(gpxString string) string {
	suffix := "    </trkseg>\n  </trk>\n</gpx>"

	preparedString := strings.TrimSuffix(gpxString, suffix)

	return preparedString
}

func PrepareSecondaryFile(gpxString string) string {
	firstElement := "    <name>Fietsen</name>\n"

	preparedString := strings.Trim(gpxString, firstElement)

	return preparedString
}
