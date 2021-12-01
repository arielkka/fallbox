package handler

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"os"
)

func getImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	return image, err
}

func convertBytesToJPG(body []byte, id string) error {
	img, err := jpeg.Decode(bytes.NewReader(body))
	if err != nil {
		return err
	}
	out, err := os.Create(fmt.Sprintf("./%s.jpg", id))
	if err != nil {
		return err
	}
	defer out.Close()
	err = jpeg.Encode(out, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func convertBytesToPNG(body []byte, id string) error {
	img, err := png.Decode(bytes.NewReader(body))
	if err != nil {
		return err
	}
	out, err := os.Create(fmt.Sprintf("./%s.png", id))
	if err != nil {
		return err
	}
	defer out.Close()
	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

func convertPNGToBytes(im image.Image) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, im)
	if err != nil {
		return nil, err
	}
	bts := buffer.Bytes()
	return bts, nil
}

func convertJPGToBytes(im image.Image) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := jpeg.Encode(buffer, im, nil)
	if err != nil {
		return nil, err
	}
	bts := buffer.Bytes()
	return bts, nil
}
