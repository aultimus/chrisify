# gofaceswap

Library providing faceswap functionality built on top of haar cascade classification implemented in opencv.

Forked from github.com/zikes/chrisify

## Linux Install

1. Install the OpenCV Developer package. On Ubuntu systems that's `sudo apt install libopencv-dev`

2. `go get -u github.com/aultimus/gofaceswap`

3. `cd $GOPATH/src/github.com/aultimus/gofaceswap && go install ./...`

## Usage

The most useful function is the FaceSwap library function, main provides example usage, which can be run as a binary like:

gofaceswap --faces ~/faces_dir --input ~/in.jpg > out.jpg

where faces_dir is a directory containing pngs of faces to be swapped
