package data

import "errors"

var (
	// ErrInvalidPageToken the token cannot be built.
	ErrInvalidPageToken = errors.New("invalid page token")
)
