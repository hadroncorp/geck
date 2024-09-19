package security

import (
	"github.com/emirpasic/gods/v2/sets"
)

type Principal interface {
	ID() string
	Sub() string
	Username() string
	Authorities() sets.Set[string]
}

type PrincipalTemplate struct {
	Identifier   string
	Subject      string
	User         string
	AuthoritySet sets.Set[string]
}

var _ Principal = (*PrincipalTemplate)(nil)

func (b PrincipalTemplate) ID() string {
	return b.Identifier
}

func (b PrincipalTemplate) Sub() string {
	return b.Subject
}

func (b PrincipalTemplate) Username() string {
	return b.User
}

func (b PrincipalTemplate) Authorities() sets.Set[string] {
	return b.AuthoritySet
}
