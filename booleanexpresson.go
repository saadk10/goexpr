package expr

// BooleanExpression is the interface that evaluates boolean expression
type BooleanExpression interface {
	evaluate(ctx EvaluationContext) bool
	GetBoolean(object interface{}) bool
	toString() string
}

// BooleanExpressionImpl implements BooleanExpression
type BooleanExpressionImpl struct {
	*ExpressionImpl
}

// NewBooleanExpression returns an instance of BooleanExpressionImpl
func NewBooleanExpression(context ExpressionContext) *BooleanExpressionImpl {
	boolExpr := &BooleanExpressionImpl{}

	boolExpr.ExpressionImpl = NewExpression(context)

	return boolExpr
}
