package encryptionfx

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/security/encryption"
)

var EncryptorAESModule = fx.Module("encryptor_aes",
	fx.Provide(
		env.ParseAs[encryption.ConfigEncryptor],
		fx.Annotate(
			encryption.NewEncryptorAES,
			fx.As(new(encryption.Encryptor)),
		),
	),
)
