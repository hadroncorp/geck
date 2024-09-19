package data

// ComparisonOperator an operator for comparisons.
type ComparisonOperator uint16

// LogicalOperator an operator for logical operations (concatenations, e.g. AND, OR).
type LogicalOperator uint8

const (
	OperatorEquals ComparisonOperator = iota + 1
	OperatorGreaterThan
	OperatorGreaterThanEquals
	OperatorLessThan
	OperatorLessThanEquals
	OperatorBetween
	OperatorNotBetween
	OperatorNotEquals
	OperatorIn
	OperatorNotIn
	OperatorLike
	OperatorNotLike
	OperatorExists
	OperatorNotExists
	OperatorIsNull
	OperatorIsNotNull
)

const (
	LogicalOperatorAnd LogicalOperator = iota + 1
	LogicalOperatorOr
)
