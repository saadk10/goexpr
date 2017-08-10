package expr

// Operator defines the interface
type Operator interface {
	getInputElements() []interface{}
	getName() string
	getReturnType() TypeImpl
	createExpression(operatorContext OperatorContext) Expression
}

// OperatorImpl implements Operator
type OperatorImpl struct {
	*IdentifierImpl
	returnType    TypeImpl
	inputElements []interface{}
}

// NewOperator returns an instance of OperatorImpl
func NewOperator(name, key string, returnType TypeImpl, inputElements []interface{}) *OperatorImpl {
	oper := &OperatorImpl{}

	oper.IdentifierImpl = NewIdentifierNameKey(name, key)
	oper.returnType = returnType
	oper.inputElements = inputElements

	return oper
}

func (o *OperatorImpl) getInputElements() []interface{} {
	return o.inputElements
}

func (o *OperatorImpl) getReturnType() TypeImpl {
	return o.returnType
}
