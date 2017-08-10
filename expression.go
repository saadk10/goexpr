package expr

// Expression defines the interface
type Expression interface {
	getType() TypeImpl
	isConstant() bool
	toString() string
	getExpressionContext() *ExpressionContext
}

// ExpressionImpl implements Expression
type ExpressionImpl struct {
	Context ExpressionContext
	Str     string
}

// NewExpression returns an instance of ExpressionImpl
func NewExpression(context ExpressionContext) *ExpressionImpl {
	expr := &ExpressionImpl{}
	expr.Context = context
	return expr
}

func (e *ExpressionImpl) setContext(ctx ExpressionContext) {
	e.Context = ctx
}

func (e *ExpressionImpl) setString(str string) {
	e.Str = str
}

func (e *ExpressionImpl) getType() TypeImpl {
	return e.Context.GetType()
}

func (e *ExpressionImpl) isConstant() bool {
	return false
}

func (e *ExpressionImpl) getLanguage() Language {
	return e.Context.getLanguage()
}

func (e *ExpressionImpl) getExpressionContext() *ExpressionContext {
	return &e.Context
}

func (e *ExpressionImpl) toString() string {
	return e.Str
}

func (e *ExpressionImpl) isList(str string) bool {
	return false
}

func (e *ExpressionImpl) evaluate() string {
	return e.Str
}
