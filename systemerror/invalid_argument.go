package systemerror

import (
	"errors"
	"fmt"
	"strings"
)

// ErrInvalidArgument the given argument is invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// NewInvalidArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
func NewInvalidArgument[T any](argumentName string, gotValue, wantValue T) SystemError {
	return SystemError{
		ErrStatus:  StatusInvalidArgument,
		ErrReason:  "INVALID_ARGUMENT",
		ErrMessage: fmt.Sprintf("argument '%s' is invalid", argumentName),
		ErrMetadata: map[string]string{
			"want_value": fmt.Sprintf("%v", wantValue),
			"got_value":  fmt.Sprintf("%v", gotValue),
		},
		StaticError: ErrInvalidArgument,
	}
}

// NewMissingArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
// Attaches 'MISSING_ARGUMENT' reason.
func NewMissingArgument(argumentName string) SystemError {
	return SystemError{
		ErrStatus:   StatusInvalidArgument,
		ErrReason:   "MISSING_ARGUMENT",
		ErrMessage:  fmt.Sprintf("argument '%s' is missing", argumentName),
		ErrMetadata: nil,
		StaticError: ErrInvalidArgument,
	}
}

// NewInvalidFormatArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
// Attaches 'INVALID_FORMAT' reason.
func NewInvalidFormatArgument(argumentName, expectedFormat string) SystemError {
	return SystemError{
		ErrStatus:  StatusInvalidArgument,
		ErrReason:  "INVALID_FORMAT",
		ErrMessage: fmt.Sprintf("'%s' has an invalid format, expected [%s]", argumentName, expectedFormat),
		ErrMetadata: map[string]string{
			"expected_format": expectedFormat,
		},
		StaticError: ErrInvalidArgument,
	}
}

// NewArgumentNotOneOf allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
// Attaches 'NOT_ONE_OF' reason.
func NewArgumentNotOneOf(argumentName string, expectedValues ...string) SystemError {
	expValuesStr := strings.Join(expectedValues, ",")
	return SystemError{
		ErrStatus:  StatusInvalidArgument,
		ErrReason:  "NOT_ONE_OF",
		ErrMessage: fmt.Sprintf("'%s' is not one of the expected values [%s]", argumentName, expValuesStr),
		ErrMetadata: map[string]string{
			"expected_values": expValuesStr,
		},
		StaticError: ErrInvalidArgument,
	}
}

// NewNotEqualsArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
// Attaches 'NOT_EQUALS' reason.
func NewNotEqualsArgument(argumentName, expectedValue string) SystemError {
	return SystemError{
		ErrStatus:  StatusInvalidArgument,
		ErrReason:  "NOT_EQUALS",
		ErrMessage: fmt.Sprintf("'%s' has an invalid value, expected [%s]", argumentName, expectedValue),
		ErrMetadata: map[string]string{
			"expected_value": expectedValue,
		},
		StaticError: ErrInvalidArgument,
	}
}

func newInvalidPrefixSuffix(argumentName, expectedType, expectedValue string) SystemError {
	return SystemError{
		ErrStatus:  StatusInvalidArgument,
		ErrReason:  "INVALID_FORMAT",
		ErrMessage: fmt.Sprintf("'%s' has an invalid format, expected %s [%s]", argumentName, expectedType, expectedValue),
		ErrMetadata: map[string]string{
			"expected_format": expectedValue,
		},
		StaticError: ErrInvalidArgument,
	}
}

// NewInvalidPrefixArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// Attaches 'INVALID_FORMAT' reason and adds 'prefix' as expected type in message.
func NewInvalidPrefixArgument(argumentName, expectedPrefix string) SystemError {
	return newInvalidPrefixSuffix(argumentName, "prefix", expectedPrefix)
}

// NewInvalidSuffixArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
// Attaches 'INVALID_FORMAT' reason and adds 'suffix' as expected type in message.
func NewInvalidSuffixArgument(argumentName, expectedSuffix string) SystemError {
	return newInvalidPrefixSuffix(argumentName, "suffix", expectedSuffix)
}

// NewInvalidNoPrefixArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
// Attaches 'INVALID_FORMAT' reason and adds 'no prefix' as expected type in message.
func NewInvalidNoPrefixArgument(argumentName, excludedPrefix string) SystemError {
	return newInvalidPrefixSuffix(argumentName, "no prefix", excludedPrefix)
}

// NewInvalidNoSuffixArgument allocates a new SystemError using StatusInvalidArgument and ErrInvalidArgument.
//
// An invalid argument has been detected.
// Attaches 'INVALID_FORMAT' reason and adds 'no suffix' as expected type in message.
func NewInvalidNoSuffixArgument(argumentName, excludedSuffix string) SystemError {
	return newInvalidPrefixSuffix(argumentName, "no suffix", excludedSuffix)
}
