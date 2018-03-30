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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	p := color.Palette{}
	palet := image.NewPaletted(r, p)
	for i := 0; i < 50; i++ {
		palet.Set(i, i, cr)
	}
	w, err = os.OpenFile("paletted.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(w, palet)
	if err != nil {
		// log.Fatal(err) // png: invalid format: bad palette length: 0
	}

	point := image.Pt(50, 50)
	log.Printf("%v\n", point)

	uniform := image.NewUniform(cr)
	log.Printf("%+v\n", uniform)
	w, err = os.OpenFile("uniform.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(w, uniform)
	if err != nil {
		log.Fatal(err)
	}
}
