package security

import (
	"context"

	"github.com/hadroncorp/geck/systemerror"
)

type PrincipalContextType string

const PrincipalContextKey PrincipalContextType = "geck.security.principal"

func GetPrincipalFromContext(ctx context.Context) (Principal, error) {
	principal, ok := ctx.Value(PrincipalContextKey).(Principal)
	if !ok {
		return nil, systemerror.NewUnauthenticated()
	}
	return principal, nil
}
