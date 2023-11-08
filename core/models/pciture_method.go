package models

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/errors"
)


func (picR *PictureRaw) ToPicture() (*Picture, error) {
	if picR.Binary == nil && !picR.Downloaded {
		return nil, errors.ErrPictureDownloadFailed
	}
	pictureDB := &Picture{
		DirectURL:  picR.DirectURL,
		Width:      picR.Width,
		Height:     picR.Height,
		Hash:       picR.Hash,
		BlurScore:  picR.BlurScore,
		Downloaded: picR.Downloaded,
	}
	format := strings.Split(picR.DirectURL, ".")[len(strings.Split(picR.DirectURL, "."))-1]
	if picR.Binary != nil {
		filePath, err := savePicture(picR.Binary, format)
		if err != nil {
			return nil, err
		}
		pictureDB.FilePath = filePath
	}
	return pictureDB, nil
}

func savePicture(binary []byte, format string) (string, error) {
	fileName := strconv.Itoa(int(time.Now().UnixMilli())) + "." + format
	prefix := config.Cfg.App.ImagePrefix
	dir := prefix + "images/" + time.Now().Format("2006/01/02")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	filePath := dir + "/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = file.Write(binary)
	if err != nil {
		return "", err
	}

	// if strings.HasPrefix(prefix, "http") {
	// 	TODO: upload to OSS
	// }

	return strings.TrimPrefix(filePath, prefix), nil
}


func (p *Picture) ToResp() *RespPicture {
	return &RespPicture{
		Status:    0,
		Message:   "success",
		CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: p.UpdatedAt.Format("2006-01-02 15:04:05"),
		DirectURL: p.DirectURL,
		Width:     p.Width,
		Height:    p.Height,
		BlurScore: p.BlurScore,
		Hash:      p.Hash,
	}
}