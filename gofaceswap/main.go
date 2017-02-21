package main

import (
	"flag"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/aultimus/gofaceswap"
)

var haarCascade = flag.String("haar", "haarcascade_frontalface_alt.xml", "The location of the Haar Cascade XML configuration to be provided to OpenCV.")
var facesDir = flag.String("faces", "", "The directory to search for faces to draw on the input image")
var inFile = flag.String("input", "", "input image to draw faces on")

func main() {
	flag.Parse()

	var facesPath string
	var err error

	if *inFile == "" {
		panic("no input file specified")
	}

	if *facesDir != "" {
		facesPath, err = filepath.Abs(*facesDir)
		if err != nil {
			panic(err)
		}
	}

	var outFaces gofaceswap.FaceList
	err = outFaces.Load(facesPath)
	if err != nil {
		panic(err)
	}
	if len(outFaces) == 0 {
		panic("no faces found")
	}

	baseImage := gofaceswap.LoadImage(*inFile)
	gofaceswap.FaceSwap(baseImage, outFaces, *haarCascade, os.Stdout)
}
