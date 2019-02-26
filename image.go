package main

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

// resize image to fix telegram profile image
func fixSizeImage(b []byte) ([]byte, error) {
	read := bytes.NewReader(b)
	img, f, err := image.Decode(read)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rec := img.Bounds()
	width := rec.Dx()
	height := rec.Dy()
	log.Printf("Resize Image %d X %d %s", width, height, f)
	if height > width {
		height = width
	} else if width > height {
		width = height
	}
	// Resize the image
	resizeImg := imaging.Fit(img, width, height, imaging.Lanczos)
	// Create a new black background image
	bgImage := imaging.New(width, height, color.Black)
	// paste the resized images into background image.
	img = imaging.PasteCenter(bgImage, resizeImg)
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	err = imaging.Encode(writer, img, imaging.JPEG)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return buf.Bytes(), nil
}
