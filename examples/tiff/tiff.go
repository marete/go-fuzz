package tiff

import (
	"bytes"
	"golang.org/x/image/tiff"
)

func Fuzz(data []byte) int {
	cfg, err := tiff.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	if cfg.Width*cfg.Height > 1e6 {
		return 0
	}
	img, err := tiff.Decode(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	var w bytes.Buffer
	err = tiff.Encode(&w, img, nil)
	if err != nil {
		panic(err)
	}
	return 1
}
