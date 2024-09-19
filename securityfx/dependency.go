package securityfx

import (
	"github.com/MicahParks/keyfunc/v3"
	"github.com/caarlos0/env/v11"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/security"
)

var GenericJWTModule = fx.Module("generic_jwt",
	fx.Provide(
		env.ParseAs[security.ConfigJWT],
	),
)

var CognitoModule = fx.Module("amazon_cognito",
	fx.Provide(
		env.ParseAs[security.ConfigCognitoKeys],
		fx.Annotate(
			security.NewAmazonCognitoKeysJWK,
			fx.As(new(keyfunc.Keyfunc)),
		),
		security.NewConfigCognitoJWT,
		fx.Annotate(
			security.NewPrincipalManagerCognito,
			fx.As(new(security.PrincipalFactory[*jwt.Token])),
		),
	),
)
