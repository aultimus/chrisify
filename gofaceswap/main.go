package main

import (
	"flag"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/aultimus/gofaceswap"
)

// usage: go run gofaceswap/main.go --haar haarcascade_frontalface_alt.xml  --faces ~/Desktop/faces --input ~/Desktop/cats-jones.jpg > out.jpg
func main() {
	var haarCascade = flag.String("haar", "haarcascade_frontalface_alt.xml", "The location of the Haar Cascade XML configuration to be provided to OpenCV.")
	var facesDir = flag.String("faces", "", "The directory to search for faces to draw on the input image")
	var inFile = flag.String("input", "", "input image to draw faces on")
	flag.Parse()

	if *inFile == "" {
		panic("no input file specified")
	}

	if *facesDir == "" {
		panic("no faces dir specified")
	}

	facesPath, err := filepath.Abs(*facesDir)
	if err != nil {
		panic(err)
	}

	outFaces, err := gofaceswap.FaceListFromDir(facesPath)
	if err != nil {
		panic(err)
	}
	if len(outFaces) == 0 {
		panic("no faces found")
	}

	baseImage := gofaceswap.LoadImage(*inFile)
	gofaceswap.FaceSwap(baseImage, outFaces, *haarCascade, os.Stdout)
}
