package errors

import (
	"errors"
)

var (
	ErrArtworkNotFound = errors.New("artwork not found")
	ErrPictureNotFound = errors.New("picture not found")
)
