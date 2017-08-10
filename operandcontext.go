package expr

// OperandContext defines the struct
type OperandContext struct {
	ExpressionContext
	fieldValues []interface{}
}

// NewOperandContext return an instance of OperandContext
func NewOperandContext(expectedType Type, fieldValues []interface{}, lang *LanguageImpl) *OperandContext {
	ctx := &OperandContext{}

	ctx.ExpressionContext.Typ = *expectedType.getType()
	ctx.fieldValues = fieldValues

	return ctx
}
