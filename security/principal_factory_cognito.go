package security

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/v2/sets/hashset"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

type PrincipalFactoryCognito struct {
}

var _ PrincipalFactory[*jwt.Token] = (*PrincipalFactoryCognito)(nil)

func NewPrincipalManagerCognito() PrincipalFactoryCognito {
	return PrincipalFactoryCognito{}
}

func (p PrincipalFactoryCognito) NewContextWithPrincipal(parent context.Context, args *jwt.Token) (context.Context, error) {
	principal, err := p.convertToPrincipal(args)
	if err != nil {
		return nil, err
	}
	return context.WithValue(parent, PrincipalContextKey, principal), nil
}

func (p PrincipalFactoryCognito) convertToPrincipal(args *jwt.Token) (Principal, error) {
	claims, ok := args.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot cast jwt claims")
	}
	sub, _ := claims.GetSubject()
	username, _ := claims["username"].(string)
	scopesRaw, _ := claims["scope"].(string)
	scopes := strings.Split(scopesRaw, " ")

	groupsRaw, _ := claims["cognito:groups"].([]any)
	groups := lo.Map(groupsRaw, func(src any, index int) string {
		return fmt.Sprintf("%v", src)
	})

	authoritySet := hashset.New[string](scopes...)
	authoritySet.Add(groups...)
	return PrincipalTemplate{
		Identifier:   username,
		Subject:      sub,
		User:         username,
		AuthoritySet: authoritySet,
	}, nil
}
