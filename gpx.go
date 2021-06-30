package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	reg := regexp.MustCompile("(.*\n)+.*<trk>\n.*</name>\n.*<trkseg>\n")

	preparedString := reg.ReplaceAllString(gpxString, "")

	return preparedString
}

func CombineGpxStrings(primary, secondary string) string {
	return primary + secondary
}

func CreateGpxFile(gpxString string, path string) {
	err := ioutil.WriteFile(path, []byte(gpxString), 0644)

	if err != nil {
		fmt.Printf("Shit got fucked %v", err)
	}
}
