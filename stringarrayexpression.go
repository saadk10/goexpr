package expr

// StringArrayExpression is the interface that evaluates string expression
type StringArrayExpression interface {
	evaluate(ctx EvaluationContext) []string
	toString() string
}

// StringArrayExpressionImpl implements StringArrayExpression
type StringArrayExpressionImpl struct {
	ListExpression
}
