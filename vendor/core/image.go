package core

import (
	"bufio"
	"fmt"
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

func Resize(source, width, height, filter string) {
	var resampleFilter imaging.ResampleFilter
	if filter == "NearestNeighbor" {
		resampleFilter = imaging.NearestNeighbor
	} else if filter == "Box" {
		resampleFilter = imaging.Box
	} else if filter == "Linear" {
		resampleFilter = imaging.Linear
	} else if filter == "Hermite" {
		resampleFilter = imaging.Hermite
	} else if filter == "MitchellNetravali" {
		resampleFilter = imaging.MitchellNetravali
	} else if filter == "CatmullRom" {
		resampleFilter = imaging.CatmullRom
	} else if filter == "BSpline" {
		resampleFilter = imaging.BSpline
	} else if filter == "Gaussian" {
		resampleFilter = imaging.Gaussian
	} else if filter == "Lanczos" {
		resampleFilter = imaging.Lanczos
	} else if filter == "Hann" {
		resampleFilter = imaging.Hann
	} else if filter == "Hamming" {
		resampleFilter = imaging.Hamming
	} else if filter == "Blackman" {
		resampleFilter = imaging.Blackman
	} else if filter == "Bartlett" {
		resampleFilter = imaging.Bartlett
	} else if filter == "Welch" {
		resampleFilter = imaging.Welch
	} else if filter == "Cosine" {
		resampleFilter = imaging.Cosine
	} else {
		fmt.Fprintf(os.Stderr, "필터가 올바르지 않습니다.")
		return
	}

	nWidth, err := strconv.Atoi(width)
	if err != nil {
		panic(err)
	}
	nHeight, err := strconv.Atoi(height)
	if err != nil {
		panic(err)
	}

	src, err := imaging.Open(source)
	if err != nil {
		panic(err)
	}

	var dst *image.NRGBA
	dst = imaging.Resize(src, nWidth, nHeight, resampleFilter)

	err = imaging.Save(dst, source)
	if err != nil {
		panic(err)
	}
}
