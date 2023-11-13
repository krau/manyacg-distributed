package errors

import (
	"errors"
)

var (
	ErrArtworkNotFound                 = errors.New("artwork not found")
	ErrPictureNotFound                 = errors.New("picture not found")
	ErrMessengerAzureNotInitialized    = errors.New("messenger azure not initialized")
	ErrMessengerRabbitMQNotInitialized = errors.New("messenger rabbitmq not initialized")
	ErrPictureDownloadFailed           = errors.New("picture download failed")
	ErrUnknownSaveType                 = errors.New("unknown save type")
)
