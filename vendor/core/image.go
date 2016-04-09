package core

import (
	"bufio"
	"github.com/disintegration/imaging"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
)

func Crop(source, target, x, y, width, height string) {
	nX, err := strconv.Atoi(x)
	nY, err := strconv.Atoi(y)
	nWidth, err := strconv.Atoi(width)
	nHeight, err := strconv.Atoi(height)
	f, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, ext, err := image.Decode(bufio.NewReader(f))
	if err != nil {
		panic(err)
	}

	crp, err := cutter.Crop(src, cutter.Config{
		Width:  nWidth,
		Height: nHeight,
		Anchor: image.Point{nX, nY},
	})

	r, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	if ext == "jpg" {
		var opt jpeg.Options
		opt.Quality = 100
		err = jpeg.Encode(r, crp, &opt)
	} else if ext == "png" {
		err = png.Encode(r, crp)
	}

	if err != nil {
		panic(err)
	}

}

func Rotate(source, direction string) {
	src, err := imaging.Open(source)
	if err != nil {
		panic(err)
	}

	var dst *image.NRGBA
	if direction == "left" {
		dst = imaging.Rotate90(src)
	} else if direction == "right" {
		dst = imaging.Rotate270(src)
	}

	err = imaging.Save(dst, source)
	if err != nil {
		panic(err)
	}
}
