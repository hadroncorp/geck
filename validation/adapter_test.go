package validation

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hadroncorp/geck/systemerror"
)

func TestWrapper(t *testing.T) {
	type stub struct {
		SomeString   string  `validate:"required"`
		SomeFloat    float64 `validate:"lte=1"`
		SomeCurrency string  `validate:"iso4217"`
		SomeOneOf    string  `validate:"oneof=foo bar baz"`
		SomeLen      []int   `validate:"len=10"`
		SomeMax      []int   `validate:"max=1"`
		SomeMin      []int   `validate:"min=1"`
		SomeGT       float64 `validate:"gt=1"`
		SomeGTEStr   string  `validate:"gte=1"`
		SomeLT       float64 `validate:"lt=1"`
		SomeEQ       string  `validate:"eq=foo"`
	}

	validator := NewGoPlaygroundValidator()
	validatorErr := validator.Validate(context.Background(), stub{
		SomeMax: []int{1, 2, 3},
		SomeLT:  2,
	})
	err := adapterGoPlaygroundErrors(validatorErr)
	t.Log(err)

	assert.True(t, errors.Is(err, systemerror.ErrInvalidArgument))
}
