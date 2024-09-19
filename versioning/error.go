package versioning

import "errors"

// ErrInvalidSemver the given version is not a valid semantic version.
var ErrInvalidSemver = errors.New("invalid semantic version")
