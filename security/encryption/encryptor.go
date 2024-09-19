package encryption

type Encryptor interface {
	Encrypt(plainText string) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
}
