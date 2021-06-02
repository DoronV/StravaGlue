package main

import (
	"fmt"
	"io/ioutil"
)

type GpxFile struct {
	GpxString []byte
}

func Read(filepath string) GpxFile {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Errorf("error reading the file, %v", err)
	}

	gpxFile := GpxFile{
		GpxString: data,
	}

	return gpxFile
}
