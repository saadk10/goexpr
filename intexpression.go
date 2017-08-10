package expr

// IntExpression defines the interface
type IntExpression interface {
	evaluate(ctx EvaluationContext) int
	toString() string
}

// IntExpressionImpl implements IntExpression
type IntExpressionImpl struct {
	*ExpressionImpl
}

// NewIntExpression returns an instance of IntExpressionImpl
func NewIntExpression(context ExpressionContext) *IntExpressionImpl {
	ie := &IntExpressionImpl{}
	ie.ExpressionImpl = NewExpression(context)
	return ie
}
