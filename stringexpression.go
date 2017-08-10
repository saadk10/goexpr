package expr

// StringExpression defines the interface
type StringExpression interface {
	evaluate(ctx EvaluationContext) string
	toString() string
}

// StringExpressionImpl implements StringExpression
type StringExpressionImpl struct {
	*ExpressionImpl
}

var (
	any = "*"
)

// NewStringExpression returns an instance of StringExpressionImpl
func NewStringExpression(context ExpressionContext) *StringExpressionImpl {
	se := &StringExpressionImpl{}

	se.ExpressionImpl = NewExpression(context)
	return se
}
