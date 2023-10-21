package models

type SourceName string

const (
	SourcePixiv SourceName = "Pixiv"
)


func (s SourceName) String() string {
	return string(s)
}