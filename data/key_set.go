package data

import "fmt"

type KeySet struct {
	Field    string
	Operator ComparisonOperator
	Value    string
}

var _ fmt.Stringer = (*KeySet)(nil)

func (k KeySet) String() string {
	return fmt.Sprintf("%s %s %s", k.Field, k.Operator, k.Value)
}
