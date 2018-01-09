package main

import (
	"bytes"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io/ioutil"
)

const colorPivot = 256

func main() {
	// 画布大小
	width, height := 500, 500

	// 属性配置
	var g = new(gif.GIF)
	for i := 1; i < 100; i++ {
		var pal = image.NewPaletted(image.Rect(0, 0, width, height), palette.WebSafe)
		for w := 0; w < width; w++ {
			for h := 0; h < height; h++ {
				pal.Set(w, h, color.RGBA{0x66, 0x33, 0xFF, 0x00})
			}
		}
		for j := 1; j < 10; j++ {
			for k := 1; k < 10; k++ {
				pal.Set(k+width/i+j, height/i+j, color.RGBA{0x00, 0x00, 0xff, 0x00})
				pal.Set(k+width/i-j, height/i-j, color.RGBA{0x00, 0x00, 0xff, 0x00})
			}
		}
		g.Image = append(g.Image, pal)
		g.Delay = append(g.Delay, i%10)
	}

	// 编码
	var w = new(bytes.Buffer)
	err := gif.EncodeAll(w, g)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("test.gif", w.Bytes(), 0755)
	if err != nil {
		panic(err)
	}
}
