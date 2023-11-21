package errors

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

var (
	ErrArtworkNotFound                 = errors.New("artwork not found")
	ErrPictureNotFound                 = errors.New("picture not found")
	ErrMessengerAzureNotInitialized    = errors.New("messenger azure not initialized")
	ErrMessengerRabbitMQNotInitialized = errors.New("messenger rabbitmq not initialized")
	ErrPictureDownloadFailed           = errors.New("picture download failed")
	ErrPictureSaveFailed               = errors.New("picture save failed")
	ErrUnknownSaveType                 = errors.New("unknown save type")
	ErrRedisKeyNotFound                = redis.Nil
	ErrWebdavClientNotInitialized      = errors.New("webdav client not initialized")
)
