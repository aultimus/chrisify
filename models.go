package gofaceswap

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/aultimus/gofaceswap/facefinder"
	"github.com/disintegration/imaging"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Face struct {
	image.Image
}

func (f *Face) LoadFile(file string) error {
	reader, err := os.Open(file)
	if err != nil {
		return err
	}
	f.Image, _, err = image.Decode(reader)
	if err != nil {
		return err
	}
	return nil
}

func NewFace(file string) (*Face, error) {
	face := &Face{}
	if err := face.LoadFile(file); err != nil {
		return face, err
	}
	return face, nil
}

func NewMustFace(file string) *Face {
	face, err := NewFace(file)
	if err != nil {
		panic(err)
	}
	return face
}

type FaceList []*Face

func (fl FaceList) Random() image.Image {
	i := rand.Intn(len(fl))
	face := fl[i]
	if rand.Intn(2) == 0 {
		return imaging.FlipH(face.Image)
	}
	return face.Image
}

func (fl *FaceList) Load(dir string) error {
	if dir == "" {
		return nil
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".png" {
			f, err := NewFace(path.Join(dir, file.Name()))
			if err != nil {
				return err
			}
			*fl = append(*fl, f)
		}
	}
	return nil
}

// FaceListFromDir searches the specified directory for pngs,
// and constructs a facelist from them.
func FaceListFromDir(dir string) (FaceList, error) {
	var fl FaceList
	err := fl.Load(dir)
	return fl, err
}

func FaceSwap(baseImage image.Image, outFaces FaceList, haarFPath string,
	out io.Writer) {

	finder := facefinder.NewFinder(haarFPath)

	faces := finder.Detect(baseImage)

	bounds := baseImage.Bounds()

	canvas := canvasFromImage(baseImage)

	for _, face := range faces {
		rect := rectMargin(30.0, face)

		newFace := outFaces.Random()
		if newFace == nil {
			panic("nil face")
		}
		chrisFace := imaging.Fit(newFace, rect.Dx(), rect.Dy(), imaging.Lanczos)

		draw.Draw(
			canvas,
			rect,
			chrisFace,
			bounds.Min,
			draw.Over,
		)
	}

	if len(faces) == 0 {
		face := imaging.Resize(
			outFaces[0],
			bounds.Dx()/3,
			0,
			imaging.Lanczos,
		)
		face_bounds := face.Bounds()
		draw.Draw(
			canvas,
			bounds,
			face,
			bounds.Min.Add(image.Pt(-bounds.Max.X/2+face_bounds.Max.X/2, -bounds.Max.Y+int(float64(face_bounds.Max.Y)/1.9))),
			draw.Over,
		)
	}

	jpeg.Encode(out, canvas, &jpeg.Options{jpeg.DefaultQuality})

}
