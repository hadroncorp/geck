package data

// CriteriaFilter a filter operation for a Criteria specification.
type CriteriaFilter struct {
	// Field name of the field to filter.
	Field string
	// Operator comparison operator.
	Operator ComparisonOperator
	// Value slice of values. Depending on the Operator, the processing component might accept
	Value []any
}

// CriteriaOrdering the ordering technique of a Criteria operation.
type CriteriaOrdering struct {
	// Field name of the field dataset will be ordered by.
	Field string
	// OrderType the type of ordering to use.
	OrderType OrderType
}

// Criteria a domain-specific language (DSL) structure used by PagingRepository instances to specify
// read operations arguments.
type Criteria struct {
	// PageSize maximum number of items to fetch.
	PageSize int64 `validate:"omitempty,min=1,max=250"`
	// PageToken
	PageToken       PageToken `validate:"omitempty,max=255"`
	Ordering        CriteriaOrdering
	LogicalOperator LogicalOperator
	Filters         []CriteriaFilter
}

type CriteriaFields map[string]string
