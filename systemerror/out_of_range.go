package systemerror

import (
	"errors"
	"fmt"
)

// ErrOutOfRange the given input/resource/data is out of range.
var ErrOutOfRange = errors.New("out of range")

// NewOutOfRange allocates a SystemError with StatusOutOfRange and ErrOutOfRange.
func NewOutOfRange(reason, name string, metadata map[string]string) SystemError {
	return SystemError{
		ErrStatus:   StatusOutOfRange,
		ErrReason:   reason,
		ErrMessage:  fmt.Sprintf("'%s' is out of range", name),
		ErrMetadata: metadata,
		StaticError: ErrOutOfRange,
	}
}

var outOfRangeTypeNameMap = map[string]string{
	"lt":  "less than",
	"lte": "less than or equals",
	"gt":  "greater than",
	"gte": "greater than or equals",
	"max": "maximum",
	"min": "minimum",
	"len": "length",
}

// NewArgumentOutOfRange allocates a SystemError with StatusOutOfRange and ErrOutOfRange.
//
// Attaches 'ARGUMENT_OUT_OF_RANGE' reason and both, n and k factors, to metadata.
func NewArgumentOutOfRange(argumentName, n, k any) SystemError {
	return SystemError{
		ErrStatus:  StatusOutOfRange,
		ErrReason:  "ARGUMENT_OUT_OF_RANGE",
		ErrMessage: fmt.Sprintf("'%s' is out of range, expected [%v,%v]", argumentName, n, k),
		ErrMetadata: map[string]string{
			"range_n": fmt.Sprintf("%v", n),
			"range_k": fmt.Sprintf("%v", k),
		},
		StaticError: ErrOutOfRange,
	}
}

// NewArgumentOutOfRangeSingle allocates a SystemError with StatusOutOfRange and ErrOutOfRange.
//
// Attaches 'ARGUMENT_OUT_OF_RANGE' reason and n range factor to metadata.
func NewArgumentOutOfRangeSingle(argumentName string, rangeType string, n any) SystemError {
	return SystemError{
		ErrStatus: StatusOutOfRange,
		ErrReason: "ARGUMENT_OUT_OF_RANGE",
		ErrMessage: fmt.Sprintf("'%s' is out of range, expected %s [%v]", argumentName,
			outOfRangeTypeNameMap[rangeType], n),
		ErrMetadata: map[string]string{
			"range": fmt.Sprintf("%v", n),
		},
		StaticError: ErrOutOfRange,
	}
}
