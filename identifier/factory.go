package identifier

import (
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/ksuid"
)

type Factory interface {
	NewIdentifier() (string, error)
}

// FactoryKSUID is the Factory implementation using segmentio's KSUID.
type FactoryKSUID struct{}

var _ Factory = (*FactoryKSUID)(nil)

func NewFactoryKSUID() FactoryKSUID {
	return FactoryKSUID{}
}

func (f FactoryKSUID) NewIdentifier() (string, error) {
	id, err := ksuid.NewRandomWithTime(time.Now().UTC())
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// FactoryUUID is the Factory implementation using UUID v7.
type FactoryUUID struct{}

var _ Factory = (*FactoryUUID)(nil)

func NewFactoryUUID() FactoryUUID {
	return FactoryUUID{}
}

func (f FactoryUUID) NewIdentifier() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
