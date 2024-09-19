package systemerror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hadroncorp/geck/systemerror"
)

func TestSomeTest(t *testing.T) {
	var err error
	err = systemerror.NewDomain("", "some error", nil)
	assert.True(t, errors.Is(err, systemerror.ErrDomain))

	err = systemerror.NewArgumentOutOfRange("items", 1.10, 99.99)
	t.Log(err)
	assert.True(t, errors.Is(err, systemerror.ErrOutOfRange))
}
