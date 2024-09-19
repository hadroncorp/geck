package systemerror

import (
	"errors"
	"strings"
)

// ErrPermissionDenied the specified principal has no permission to access a resource.
var ErrPermissionDenied = errors.New("permission denied")

// NewPermissionDeniedAuthorities Allocates a SystemError with StatusPermissionDenied and ErrPermissionDenied.
//
// Principal has no authority to perform action.
// Sets 'PRINCIPAL_MISSING_AUTHORITIES' as reason.
func NewPermissionDeniedAuthorities(principalName string, authorities, expAuthorities []string) SystemError {
	return SystemError{
		ErrStatus:  StatusPermissionDenied,
		ErrReason:  "PRINCIPAL_MISSING_AUTHORITIES",
		ErrMessage: "principal is not authorized to perform this operation",
		ErrMetadata: map[string]string{
			"principal":            principalName,
			"authorities":          strings.Join(authorities, " "),
			"expected_authorities": strings.Join(expAuthorities, " "),
		},
		StaticError: ErrPermissionDenied,
	}
}

// NewPermissionDeniedInvalidOwner allocates a SystemError with StatusPermissionDenied and ErrPermissionDenied.
//
// Principal is not the owner of the resource.
// Sets 'PRINCIPAL_NOT_OWNER' as reason.
func NewPermissionDeniedInvalidOwner(principalID string, expOwnerID string) SystemError {
	return SystemError{
		ErrStatus:  StatusPermissionDenied,
		ErrReason:  "PRINCIPAL_NOT_OWNER",
		ErrMessage: "principal is not authorized to perform this operation",
		ErrMetadata: map[string]string{
			"principal":      principalID,
			"expected_owner": expOwnerID,
		},
		StaticError: ErrPermissionDenied,
	}
}
