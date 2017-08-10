package expr

// False is the implementation of the operator 'False'
type False struct {
	*OperatorImpl
}

// NewFalse creates a new instance of False
func NewFalse() *False {
	falseOperator := &False{}
	falseOperator.OperatorImpl = NewOperator("false", "", booleanType, []interface{}{"false"})
	return falseOperator
}

func (o *False) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *False) getName() string {
	return o.OperatorImpl.name
}

func (o *False) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *False) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'false' operator")
	return newFalseInstance(ctx)

}

type falseInstance struct {
	*BooleanExpressionImpl
}

func newFalseInstance(ctx OperatorContext) Expression {
	falseInst := &falseInstance{}

	falseInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)

	return falseInst
}

func (i *falseInstance) evaluate(ctx EvaluationContext) bool {
	return false
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *falseInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *falseInstance) toString() string {
	return "false"
}
