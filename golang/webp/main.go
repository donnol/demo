package main

import (
	"image"
	"image/color"
	"log"
	"os"

	webpbin "github.com/nickalie/go-webpbin"
)

func main() {
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.webp")
	if err != nil {
		log.Fatal(err)
	}

	if err := webpbin.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
