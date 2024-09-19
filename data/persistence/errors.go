package persistence

import "errors"

var (
	ErrTxContextNotFound = errors.New("transaction context not found")
)
