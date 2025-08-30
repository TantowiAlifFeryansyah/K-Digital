package errors

import "errors"

var (
	ErrFileNotFound  = errors.New("file not found")
	ErrInvalidImage  = errors.New("invalid image format")
	ErrNoBorderFound = errors.New("no border detected in the image")
	ErrProcessFailed = errors.New("failed to process image")
)
