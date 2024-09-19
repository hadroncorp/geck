package security

import "github.com/caarlos0/env/v11"

func NewConfigCognitoJWT() (ConfigJWT, error) {
	cfg, err := env.ParseAs[ConfigJWT]()
	if err != nil {
		return ConfigJWT{}, err
	}
	cfg.SigningMethod = "RS256"
	return ConfigJWT{}, nil
}
