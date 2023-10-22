package processor

import (
	"bytes"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
)

func getBlurScore(picture *models.PictureRaw) {
	if picture.Binary == nil {
		return
	}
	r := bytes.NewReader(picture.Binary)
	img, _, err := image.Decode(r)
	if err != nil {
		logger.L.Errorf("Failed to decode picture: %s", err)
		return
	}
	// 转换图像为灰度图
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	laplaceImg := laplacianTransform(grayImg)
	variance := calculateVariance(laplaceImg)
	picture.BlurScore = variance
}

// 计算图像的拉普拉斯变换
func laplacianTransform(img *image.Gray) *image.Gray {
	bounds := img.Bounds()
	laplaceImg := image.NewGray(bounds)
	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			sum := int(img.GrayAt(x, y).Y*9 - img.GrayAt(x+1, y).Y - img.GrayAt(x-1, y).Y - img.GrayAt(x, y+1).Y - img.GrayAt(x, y-1).Y - img.GrayAt(x+1, y+1).Y - img.GrayAt(x-1, y+1).Y - img.GrayAt(x+1, y-1).Y - img.GrayAt(x-1, y-1).Y)
			laplaceImg.SetGray(x, y, color.Gray{uint8(sum / 8)})
		}
	}
	return laplaceImg
}

// 计算图像的方差
func calculateVariance(img *image.Gray) float64 {
	bounds := img.Bounds()
	mean := 0.0
	variance := 0.0
	pixelCount := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayValue := float64(img.GrayAt(x, y).Y)
			mean += grayValue
			pixelCount++
		}
	}

	mean /= float64(pixelCount)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayValue := float64(img.GrayAt(x, y).Y)
			variance += (grayValue - mean) * (grayValue - mean)
		}
	}

	variance /= float64(pixelCount)

	return variance
}
