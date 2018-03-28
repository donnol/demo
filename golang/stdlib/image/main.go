package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
)

type myImage struct{}

func (m myImage) ColorModel() (cm color.Model) {
	return
}

func (m myImage) Bounds() (rect image.Rectangle) {
	return
}

func (m myImage) At(x, y int) (cc color.Color) {
	return
}

func main() {
	var _ image.Image = myImage{}

	image.RegisterFormat("my image", "magic", func(io.Reader) (image.Image, error) {
		return myImage{}, nil
	}, func(io.Reader) (ic image.Config, err error) { return })

	var _ image.Image = &image.Alpha{}
	r := image.Rect(0, 0, 100, 100)
	alpha := image.NewAlpha(r)
	cr := color.Gray{}
	for i := 0; i < 50; i++ {
		alpha.Set(i, i, cr)
	}

	w, err := os.OpenFile("alpha.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(w, alpha)
	if err != nil {
		log.Fatal(err)
	}
}
