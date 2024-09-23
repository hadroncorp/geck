package data

import (
	"encoding"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/hadroncorp/geck/security/encryption"
)

const (
	pageTokenSeparator = "#"
	// PageTokenDefaultEncryptionKey the default secret key for PageToken encryption.
	PageTokenDefaultEncryptionKey = "Some_Page_Token_Key_1927_!@#$*~<" // 32 bytes, therefore, AES 256-bit
)

// PageToken Tokens are first encrypted so anybody is able to see internal system implementation details.
// Then, the token is encoded in hex format, so it can be transferred through network protocols, and thus, systems.
//
// The token is able to use different pagination mechanism.
//
// The nomenclature proposal would be this: QUERY_TYPE#NEXT_QUERY
// Nomenclature examples:
//
//   - OFFSET#100
//   - KEY_SET#name>'Foo'
//   - CURSOR#abc-foo
type PageToken []byte

var _ fmt.Stringer = PageToken{}

// this ensures token uses hex encoding rather than base64 used by json marshaller
var _ encoding.TextMarshaler = PageToken{}

// NewPageToken allocates a new PageToken instance.
func NewPageToken(encryptor encryption.Encryptor, queryType PaginationType, value string) (PageToken, error) {
	rawValue := string(queryType) + pageTokenSeparator + value
	ciphertext, err := encryptor.Encrypt(rawValue)
	if err != nil {
		return nil, err
	}
	encodedValue := make([]byte, hex.EncodedLen(len(ciphertext)))
	hex.Encode(encodedValue, ciphertext)
	return encodedValue, nil
}

// NewPageTokenOffset allocates a new PageToken instance using offset PaginationType.
// If the given value is less than zero, then PageToken will be nil.
func NewPageTokenOffset(encryptor encryption.Encryptor, value int) (PageToken, error) {
	if value < 0 {
		return nil, nil
	}
	return NewPageToken(encryptor, PaginationTypeOffset, strconv.Itoa(value))
}

// NewPageTokenKeySet allocates a new PageToken instance using key-set PaginationType.
// If the given value is less than zero, then PageToken will be nil.
func NewPageTokenKeySet(encryptor encryption.Encryptor, set KeySet) (PageToken, error) {
	return NewPageToken(encryptor, PaginationTypeKeySet, set.String())
}

// Read decomposes encrypted token to a set of PaginationType and its value.
func (p PageToken) Read(encryptor encryption.Encryptor) (string, string, error) {
	ciphertextBytes := make([]byte, hex.DecodedLen(len(p)))
	if _, err := hex.Decode(ciphertextBytes, p); err != nil {
		return "", "", nil
	}

	decryptedToken, err := encryptor.Decrypt(ciphertextBytes)
	if err != nil {
		return "", "", err
	}

	splitToken := strings.SplitN(string(decryptedToken), pageTokenSeparator, 2)
	if len(splitToken) != 2 {
		return "", "", ErrInvalidPageToken
	}
	return splitToken[0], splitToken[1], nil
}

// String retrieves encoded token.
func (p PageToken) String() string {
	return string(p)
}

// MarshalText encodes the receiver into UTF-8-encoded text and returns the result.
func (p PageToken) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}
