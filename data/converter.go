package data

import (
	"strconv"
	"strings"

	"github.com/hadroncorp/geck/internal/converter"
	"github.com/hadroncorp/geck/security/encryption"
)

// ConvertPage converts a page populated with type A to a page populated with type B.
func ConvertPage[A, B any](src Page[A], convertFunc converter.ConvertFunc[A, B]) Page[B] {
	return Page[B]{
		PreviousPageToken: src.PreviousPageToken,
		NextPageToken:     src.NextPageToken,
		TotalItems:        src.TotalItems,
		Items:             converter.ConvertMany(src.Items, convertFunc),
	}
}

// ConvertOffsetSafe converts a PageToken using an offset PaginationType.
// Returns '0' if an error was found or PaginationType is not offset.
func ConvertOffsetSafe(token PageToken, encryptor encryption.Encryptor) int {
	if len(token) == 0 {
		return 0
	}
	tokenType, valRaw, err := token.Read(encryptor)
	if err != nil || tokenType != string(PaginationTypeOffset) {
		return 0
	}
	val, _ := strconv.Atoi(valRaw)
	return val
}

// ConvertCursorSafe converts a PageToken using a cursor PaginationType.
// Returns empty string if an error was found or PaginationType is not cursor.
func ConvertCursorSafe(token PageToken, encryptor encryption.Encryptor) string {
	if len(token) == 0 {
		return ""
	}
	tokenType, valRaw, err := token.Read(encryptor)
	if err != nil || tokenType != string(PaginationTypeCursor) {
		return ""
	}
	return valRaw
}

// ConvertKeySetSafe converts a PageToken using a key-set PaginationType.
// Returns empty string if an error was found or PaginationType is not key-set.
func ConvertKeySetSafe(token PageToken, encryptor encryption.Encryptor) KeySet {
	if len(token) == 0 {
		return KeySet{}
	}
	tokenType, valRaw, err := token.Read(encryptor)
	if err != nil || tokenType != string(PaginationTypeKeySet) {
		return KeySet{}
	}
	values := strings.SplitAfterN(valRaw, " ", 3)
	if len(values) != 3 {
		return KeySet{}
	}
	return KeySet{
		Field:    values[0],
		Operator: comparisonOperatorValues[values[1]],
		Value:    values[2],
	}
}
