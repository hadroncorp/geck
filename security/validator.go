package security

import (
	"context"

	"github.com/samber/lo"

	"github.com/hadroncorp/geck/systemerror"
)

func hasAnyAuthorities(principal Principal, authorities []string) bool {
	return lo.ContainsBy(authorities, func(item string) bool {
		return principal.Authorities().Contains(item)
	})
}

func HasAnyAuthorities(ctx context.Context, authorities []string) error {
	principal, err := GetPrincipalFromContext(ctx)
	if err != nil {
		return err
	}

	isAuthZ := hasAnyAuthorities(principal, authorities)
	if !isAuthZ {
		return systemerror.NewPermissionDeniedAuthorities(principal.ID(), principal.Authorities().Values(), authorities)
	}
	return nil
}

func HasAuthorities(ctx context.Context, authorities []string) error {
	principal, err := GetPrincipalFromContext(ctx)
	if err != nil {
		return err
	}

	isAuthZ := principal.Authorities().Contains(authorities...)
	if !isAuthZ {
		return systemerror.NewPermissionDeniedAuthorities(principal.ID(), principal.Authorities().Values(), authorities)
	}
	return nil
}

func isResourceOwner(principal Principal, resourceReqOwner string) bool {
	return principal.ID() == resourceReqOwner
}

func IsResourceOwner(ctx context.Context, resourceReqOwner string) error {
	principal, err := GetPrincipalFromContext(ctx)
	if err != nil {
		return err
	}

	if isAuthZ := isResourceOwner(principal, resourceReqOwner); !isAuthZ {
		return systemerror.NewPermissionDeniedInvalidOwner(principal.Username(), resourceReqOwner)
	}
	return nil
}

func IsResourceOwnerOrHasAnyAuthorities(ctx context.Context, resourceReqOwner string, authorities []string) error {
	principal, err := GetPrincipalFromContext(ctx)
	if err != nil {
		return err
	}

	isAuthZ := hasAnyAuthorities(principal, authorities)
	if isAuthZ {
		return nil
	}

	if isAuthZ = isResourceOwner(principal, resourceReqOwner); !isAuthZ {
		return systemerror.NewPermissionDeniedInvalidOwner(principal.Username(), resourceReqOwner)
	}
	return nil
}
