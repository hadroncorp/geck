package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type EncryptorAES struct {
	SecretKey string
}

var _ Encryptor = (*EncryptorAES)(nil)

func NewEncryptorAES(cfg ConfigEncryptor) EncryptorAES {
	return EncryptorAES{
		SecretKey: cfg.SecretKey,
	}
}

func (e EncryptorAES) Encrypt(plainText string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(e.SecretKey))
	if err != nil {
		return nil, err
	}

	// IV needs to be unique, but not secret
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plainText))
	return ciphertext, nil
}

func (e EncryptorAES) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(e.SecretKey))
	if err != nil {
		return nil, err
	}

	// IV needs to be unique, but not secret
	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, nil
}
