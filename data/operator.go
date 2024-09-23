package data

import "fmt"

// ComparisonOperator an operator for comparisons.
type ComparisonOperator uint16

var _ fmt.Stringer = ComparisonOperator(0)

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

var comparisonOperatorNames = map[ComparisonOperator]string{
	OperatorEquals:            "=",
	OperatorGreaterThan:       ">",
	OperatorGreaterThanEquals: ">=",
	OperatorLessThan:          "<",
	OperatorLessThanEquals:    "<=",
	OperatorBetween:           "BETWEEN",
	OperatorNotBetween:        "NOT BETWEEN",
	OperatorNotEquals:         "NOT EQUALS",
	OperatorIn:                "IN",
	OperatorNotIn:             "NOT IN",
	OperatorLike:              "LIKE",
	OperatorNotLike:           "NOT LIKE",
	OperatorExists:            "EXISTS",
	OperatorNotExists:         "NOT_EXISTS",
	OperatorIsNull:            "IS NULL",
	OperatorIsNotNull:         "IS NOT NULL",
}

var comparisonOperatorValues = map[string]ComparisonOperator{
	"=":           OperatorEquals,
	">":           OperatorGreaterThan,
	">=":          OperatorGreaterThanEquals,
	"<":           OperatorLessThan,
	"<=":          OperatorLessThanEquals,
	"BETWEEN":     OperatorBetween,
	"NOT BETWEEN": OperatorNotBetween,
	"NOT EQUALS":  OperatorNotEquals,
	"IN":          OperatorIn,
	"NOT IN":      OperatorNotIn,
	"LIKE":        OperatorLike,
	"NOT LIKE":    OperatorNotLike,
	"EXISTS":      OperatorExists,
	"NOT EXISTS":  OperatorNotExists,
	"IS NULL":     OperatorIsNull,
	"IS NOT NULL": OperatorIsNotNull,
}

func (c ComparisonOperator) String() string {
	return comparisonOperatorNames[c]
}

const (
	LogicalOperatorAnd LogicalOperator = iota + 1
	LogicalOperatorOr
)
