package expr

import "strconv"

// IntConstant defines the struct
type IntConstant struct {
	*IntExpressionImpl
	value int
}

// NewIntConstant returns an instance of IntConstant
func NewIntConstant(value int) *IntConstant {
	ic := &IntConstant{}
	ic.value = value

	ctx := ExpressionContext{
		Typ: intType,
	}

	ic.IntExpressionImpl = NewIntExpression(ctx)

	return ic
}

func (i *IntConstant) evaluate(ctx EvaluationContext) int {
	return i.value
}

func (i *IntConstant) toString() string {
	return strconv.Itoa(i.value)
}
