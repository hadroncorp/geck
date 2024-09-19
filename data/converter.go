package data

import (
	"github.com/hadroncorp/geck/internal/converter"
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
