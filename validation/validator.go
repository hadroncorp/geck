package validation

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(ctx context.Context, v any) error
}

func ValidateMany[T any](ctx context.Context, validator Validator, items []T) error {
	for _, val := range items {
		if err := validator.Validate(ctx, val); err != nil {
			return err
		}
	}
	return nil
}

type GoPlaygroundValidator struct {
	validator *validator.Validate
}

var _ Validator = (*GoPlaygroundValidator)(nil)

func NewGoPlaygroundValidator() GoPlaygroundValidator {
	v := validator.New()
	_ = v.RegisterValidation("date", validateDate)
	return GoPlaygroundValidator{
		validator: v,
	}
}

func (g GoPlaygroundValidator) Validate(ctx context.Context, v any) error {
	return adapterGoPlaygroundErrors(g.validator.StructCtx(ctx, v))
}

func validateDate(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	_, err := time.Parse(time.DateOnly, val)
	return err == nil
}
