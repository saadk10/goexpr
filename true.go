package expr

// True is the implementation of the operator 'True'
type True struct {
	*OperatorImpl
}

// NewTrue creates a new instance of True
func NewTrue() *True {
	trueOperator := &True{}
	trueOperator.OperatorImpl = NewOperator("True", "EXPR.Operator.True", booleanType, []interface{}{"TRUE"})
	return trueOperator
}

func (o *True) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *True) getName() string {
	return o.OperatorImpl.name
}

func (o *True) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *True) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'true' operator")
	return newTrueInstance(ctx)

}

type trueInstance struct {
	*BooleanExpressionImpl
}

func newTrueInstance(ctx OperatorContext) Expression {
	trueInst := &trueInstance{}

	trueInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)

	return trueInst
}

func (i *trueInstance) evaluate(ctx EvaluationContext) bool {
	return true
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *trueInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *trueInstance) toString() string {
	return "true"
}
