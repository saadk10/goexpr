package expr

import (
	"fmt"
)

// StringLiteral defines the implementation
type StringLiteral struct {
	StringExpressionImpl
	value string
}

// NewStringLiteral returns an instance of StringLiteral
func NewStringLiteral(value string) *StringLiteral {
	sl := &StringLiteral{}

	ec := NewExpressionContext(stringType)
	sl.StringExpressionImpl = *NewStringExpression(*ec)
	sl.value = value

	return sl
}

func (sl *StringLiteral) getName() string {
	return sl.value
}

func (sl *StringLiteral) evaluate(ctx EvaluationContext) string {
	return sl.value
}

func (sl *StringLiteral) toString() string {
	return fmt.Sprintf("'%s'", sl.value)
}

func (sl *StringLiteral) getExpressionContext() *ExpressionContext {
	return &sl.StringExpressionImpl.ExpressionImpl.Context
}
