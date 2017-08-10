package expr

// ListExpression defines the struct
type ListExpression struct {
	ExpressionImpl
	exprs []Expression
}

// NewListExpression returns an instance of ListExpression
func NewListExpression(context ExpressionContext) *ListExpression {
	le := &ListExpression{}

	le.ExpressionImpl.Context = context

	return le
}

func (le *ListExpression) setExprs(exprs []Expression) {
	le.exprs = exprs
}

func (le *ListExpression) getExprs() []Expression {
	return le.exprs
}

func (le *ListExpression) toString() string {
	return "Need to Implement"
}
