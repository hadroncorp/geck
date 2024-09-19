package security

type ConfigCognitoKeys struct {
	Region     string `env:"COGNITO_REGION" envDefault:"us-east-1"`
	UserPoolID string `env:"COGNITO_USER_POOL_ID"`
}

type ConfigJWT struct {
	SigningMethod string            `env:"JWT_SIGNING_METHOD" envDefault:"HS256"`
	SigningKey    string            `env:"JWT_SIGNING_KEY,unset"`
	SigningKeys   map[string]string `env:"JWT_SIGNING_KEYS,unset"`
}
