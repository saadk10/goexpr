package expr

// IntArrayExpression defines the interfaces
type IntArrayExpression interface {
	evaluate(ctx EvaluationContext) []int
	toString() string
}

// IntArrayExpressionImpl implements IntArrayExpressions
type IntArrayExpressionImpl struct {
	*ListExpression
	vals []int
}

// NewIntArrayExpression returns an instance of NewIntArrayExpression
func NewIntArrayExpression(context ExpressionContext) *IntArrayExpressionImpl {
	intArrayExpr := &IntArrayExpressionImpl{}

	intArrayExpr.ListExpression = NewListExpression(context)
	return intArrayExpr
}
