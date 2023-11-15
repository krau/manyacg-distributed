package utils

import (
	"bytes"

	"github.com/disintegration/imaging"
)

func ResizePicture(imgByte []byte, width, height int) ([]byte, error) {
	if (0 < width && width <= 4000) || (0 < height && height <= 4000) {
		buf := bytes.NewBuffer(imgByte)
		img, err := imaging.Decode(buf)
		if err != nil {
			return nil, err
		}
		img = imaging.Resize(img, width, height, imaging.Lanczos)
		buf.Reset()
		err = imaging.Encode(buf, img, imaging.JPEG)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	return imgByte, nil
}
