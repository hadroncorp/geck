package security

import (
	"context"
	"time"

	"github.com/MicahParks/keyfunc/v3"
)

func NewAmazonCognitoKeysJWK(cfg ConfigCognitoKeys) (keyfunc.Keyfunc, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	return keyfunc.NewDefaultCtx(ctx,
		[]string{
			"https://cognito-idp." + cfg.Region + ".amazonaws.com/" + cfg.UserPoolID + "/.well-known/jwks.json",
		})
}
