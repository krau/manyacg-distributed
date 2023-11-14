package model

type PictureRaw struct {
	DirectURL  string
	Width      uint
	Height     uint
	Hash       string
	Format     string
	Binary     []byte
	BlurScore  float64
	FilePath   string
	Downloaded bool
}
