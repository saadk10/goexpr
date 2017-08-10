package expr

// OperatorContext ...
type OperatorContext struct {
	ExpressionContext
	op     Operator
	inputs []interface{}
}

// NewOperatorContext returns an instance OperatorContext
func NewOperatorContext(op Operator, inputs []interface{}, typ TypeImpl, lang *LanguageImpl) *OperatorContext {
	oc := &OperatorContext{}

	oc.inputs = inputs
	oc.op = op
	oc.Typ = typ
	oc.Lang = lang

	return oc
}

func (o *OperatorContext) getInputs() []interface{} {
	return o.inputs
}
