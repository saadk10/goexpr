package expr

import "fmt"

// And is the implementation of the operator 'and'
type And struct {
	*OperatorImpl
}

// NewAnd creates a new instance of And
func NewAnd() *And {
	and := &And{}
	and.OperatorImpl = NewOperator("And", "EXPR.Operator.And", booleanType, []interface{}{booleanType, "AND", booleanType})
	return and
}

func (o *And) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *And) getName() string {
	return o.OperatorImpl.name
}

func (o *And) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *And) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'bool and bool' operator")
	return newAndInstance(ctx)

}

type andInstance struct {
	*BooleanExpressionImpl
	expr1 BooleanExpression
	expr2 BooleanExpression
}

func newAndInstance(ctx OperatorContext) Expression {
	andInst := &andInstance{}

	andInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	andInst.expr1 = ctx.getInputs()[0].(BooleanExpression)
	andInst.expr2 = ctx.getInputs()[1].(BooleanExpression)

	return andInst
}

func (i *andInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) && i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *andInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *andInstance) toString() string {
	return fmt.Sprintf("%+v AND %+v", i.expr1.toString(), i.expr2.toString())
}
