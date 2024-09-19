package converter

import "context"

// ConvertFunc a routine used to convert A to B.
type ConvertFunc[A, B any] func(src A) B

// ConvertFuncWithContext a routine used to convert A to B taking a context.Context.
type ConvertFuncWithContext[A, B any] func(ctx context.Context, src A) B

// ConvertSafeFunc a routine used to convert A to B safely, meaning an error might be returned
// if conversion failed.
type ConvertSafeFunc[A, B any] func(src A) (B, error)

// ConvertMany converts a slice of A to a slice of B using the specified ConvertFunc routine.
func ConvertMany[A, B any](src []A, convertFunc ConvertFunc[A, B]) []B {
	buf := make([]B, 0, len(src))
	for _, a := range src {
		buf = append(buf, convertFunc(a))
	}
	return buf
}

// ConvertManyWithContext converts a slice of A to a slice of B using the specified ConvertFunc routine.
func ConvertManyWithContext[A, B any](ctx context.Context, src []A, convertFunc ConvertFuncWithContext[A, B]) []B {
	buf := make([]B, 0, len(src))
	for _, a := range src {
		buf = append(buf, convertFunc(ctx, a))
	}
	return buf
}

// ConvertManySafe converts a slice of A to a slice of B using the specified ConvertSafeFunc routine.
// If one of the conversions failed, an error will be returned and B slice will be empty.
func ConvertManySafe[A, B any](src []A, convertFunc ConvertSafeFunc[A, B]) ([]B, error) {
	buf := make([]B, 0, len(src))
	for _, a := range src {
		out, err := convertFunc(a)
		if err != nil {
			return nil, err
		}
		buf = append(buf, out)
	}
	return buf, nil
}

// ConvertNonEmptyToPtr converts A to B if not empty. Otherwise, returns nil.
func ConvertNonEmptyToPtr[A comparable, B any](src A, convertFunc ConvertFunc[A, B]) *B {
	var zeroVal A
	if src == zeroVal {
		return nil
	}
	out := convertFunc(src)
	return &out
}

// ConvertNonEmptySafe converts A to B if not empty. Otherwise, returns nil.
func ConvertNonEmptySafe[A comparable, B any](src A, convertFunc ConvertSafeFunc[A, B]) (val B, err error) {
	var zeroVal A
	if src == zeroVal {
		return
	}
	val, err = convertFunc(src)
	return
}

// ConvertZeroValuerToPtr converts A to B if not empty. Otherwise, returns nil.
func ConvertZeroValuerToPtr[A ZeroValuer, B any](src A, convertFunc ConvertFunc[A, B]) *B {
	if src.IsZero() {
		return nil
	}
	out := convertFunc(src)
	return &out
}
