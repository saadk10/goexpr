package expr

// BooleanArrayExpression is the interface that evaluates a boolean array expression
type BooleanArrayExpression interface {
	evaluate(ctx EvaluationContext) []bool
}

// BooleanArrayExpressionImpl implements BooleanArrayExpression
type BooleanArrayExpressionImpl struct {
	*ListExpression
}

// NewBooleanArrayExpression returns an instance of BooleanArrayExpressionImpl
func NewBooleanArrayExpression(context ExpressionContext) *BooleanArrayExpressionImpl {
	boolArrayExp := &BooleanArrayExpressionImpl{}
	boolArrayExp.ListExpression = NewListExpression(context)

	return boolArrayExp
}
