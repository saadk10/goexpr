package expr

// ExpressionContext defines the struct
type ExpressionContext struct {
	Typ         TypeImpl
	Lang        Language
	FieldValues []interface{}
}

// NewExpressionContext returns an instance of ExpressionContext
func NewExpressionContext(typ TypeImpl) *ExpressionContext {
	return &ExpressionContext{
		Typ: typ,
	}
}

// GetType returns the type
func (ec ExpressionContext) GetType() TypeImpl {
	return ec.Typ
}

func (ec *ExpressionContext) setType(typ TypeImpl) {
	ec.Typ = typ
}

func (ec *ExpressionContext) getLanguage() Language {
	return ec.Lang
}

func (ec *ExpressionContext) getFieldValues() []interface{} {
	return ec.FieldValues
}
