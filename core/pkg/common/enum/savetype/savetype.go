package savetype

type SaveType string

const (
	SaveTypeLocal  SaveType = "local"
	SaveTypeWebdav SaveType = "webdav"
)

func (s SaveType) String() string {
	return string(s)
}
