package versioning

import (
	"encoding"
	"fmt"

	"golang.org/x/mod/semver"
)

// SemanticVersion informational structure containing semantic version metadata.
//
// Implements fmt.Stringer, encoding.TextMarshaler and encoding.TextUnmarshaler interfaces.
type SemanticVersion struct {
	// RawValue string value used to parse semver structure.
	RawValue string
	// Major the major version (e.g. v1, v15).
	Major string
	// MajorMinor major version along minor version as well (e.g. v1.5, v3.5, v66.120).
	MajorMinor string
	// Canonical the canonical formatting of the semantic version v. It fills in any missing .MINOR or .PATCH and discards build metadata.
	Canonical string
	// Build the build suffix of the semantic version v. For example, Build("v2.1.0+meta") == "+meta". If v is an invalid
	// semantic version string, Build will be an empty string.
	Build string
	// Prerelease the prerelease suffix of the semantic version v. For example, Prerelease("v2.1.0-pre+meta") == "-pre".
	// If v is an invalid semantic version string, Prerelease will be an empty string.
	Prerelease string
}

var (
	_ fmt.Stringer             = (*SemanticVersion)(nil)
	_ encoding.TextMarshaler   = (*SemanticVersion)(nil)
	_ encoding.TextUnmarshaler = (*SemanticVersion)(nil)
)

// NewSemanticVersion parses and allocates a new SemanticVersion.
func NewSemanticVersion(rawValue string) (SemanticVersion, error) {
	if isValid := semver.IsValid(rawValue); !isValid {
		return SemanticVersion{}, ErrInvalidSemver
	}
	return SemanticVersion{
		RawValue:   rawValue,
		Major:      semver.Major(rawValue),
		MajorMinor: semver.MajorMinor(rawValue),
		Canonical:  semver.Canonical(rawValue),
		Build:      semver.Build(rawValue),
		Prerelease: semver.Prerelease(rawValue),
	}, nil
}

func (s SemanticVersion) String() string {
	return s.RawValue
}

func (s SemanticVersion) MarshalText() (text []byte, err error) {
	return []byte(s.RawValue), nil
}

func (s *SemanticVersion) UnmarshalText(text []byte) error {
	rawValue := string(text)
	if isValid := semver.IsValid(rawValue); !isValid {
		return ErrInvalidSemver
	}
	s.RawValue = rawValue
	s.Major = semver.Major(rawValue)
	s.MajorMinor = semver.MajorMinor(rawValue)
	s.Canonical = semver.Canonical(rawValue)
	s.Build = semver.Build(rawValue)
	s.Prerelease = semver.Prerelease(rawValue)
	return nil
}
