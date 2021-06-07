package main

import (
	"fmt"
	"io/ioutil"
)

type GpxFile struct {
	GpxString string
}

func Read(filepath string) GpxFile {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Errorf("error reading the file, %v", err)
	}

	dataString := string(data)

	gpxFile := GpxFile{
		GpxString: dataString,
	}

	return gpxFile
}
