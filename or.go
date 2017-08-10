package expr

import "fmt"

// Or is the implementation of the operator 'Or'
type Or struct {
	*OperatorImpl
}

// NewOr creates a new instance of Or
func NewOr() *Or {
	or := &Or{}
	or.OperatorImpl = NewOperator("OR", "EXPR.Operator.OR", booleanType, []interface{}{booleanType, "OR", booleanType})
	return or
}

func (o *Or) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *Or) getName() string {
	return o.OperatorImpl.name
}

func (o *Or) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *Or) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'bool or bool' operator")
	return newOrInstance(ctx)

}

type orInstance struct {
	*BooleanExpressionImpl
	expr1 BooleanExpression
	expr2 BooleanExpression
}

func newOrInstance(ctx OperatorContext) Expression {
	orInst := &orInstance{}

	orInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	orInst.expr1 = ctx.getInputs()[0].(BooleanExpression)
	orInst.expr2 = ctx.getInputs()[1].(BooleanExpression)

	return orInst
}

func (i *orInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) || i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *orInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *orInstance) toString() string {
	return fmt.Sprintf("%+v OR %+v", i.expr1.toString(), i.expr2.toString())
}
