package data

// OrderType the ordering type for a persistence operation (DESC, ASC).
type OrderType uint8

const (
	// OrderTypeAscending Ascending ordering type.
	OrderTypeAscending OrderType = iota + 1
	// OrderTypeDescending Descending ordering type.
	OrderTypeDescending
)
