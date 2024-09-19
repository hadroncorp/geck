package encryption

type ConfigEncryptor struct {
	SecretKey string `env:"ENCRYPTOR_SECRET_KEY,unset" envDefault:"Some_Page_Token_Key_1927_!@#$*~<"`
}
