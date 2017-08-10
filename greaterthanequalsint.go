package expr

import "fmt"

// GreaterThanEqualsInt is the implementation of the operator 'GreaterThanEqualsInt'
type GreaterThanEqualsInt struct {
	*OperatorImpl
}

// NewGreaterThanEqualsInt creates a new instance of GreaterThanEqualsInt
func NewGreaterThanEqualsInt() *GreaterThanEqualsInt {
	gte := &GreaterThanEqualsInt{}
	gte.OperatorImpl = NewOperator(">=", "EXPR.Operator.GreaterThanEquals", booleanType, []interface{}{intType, ">=", intType})
	return gte
}

func (o *GreaterThanEqualsInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *GreaterThanEqualsInt) getName() string {
	return o.OperatorImpl.name
}

func (o *GreaterThanEqualsInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *GreaterThanEqualsInt) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'int >= int' operator")
	return newGreaterThanEqualsIntInstance(ctx)

}

type greaterThanEqualsIntInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 IntExpression
}

func newGreaterThanEqualsIntInstance(ctx OperatorContext) Expression {
	gteInst := &greaterThanEqualsIntInstance{}

	gteInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)

	gteInst.expr1 = ctx.getInputs()[0].(IntExpression)
	gteInst.expr2 = ctx.getInputs()[1].(IntExpression)

	return gteInst
}

func (i *greaterThanEqualsIntInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) >= i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *greaterThanEqualsIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *greaterThanEqualsIntInstance) toString() string {
	return fmt.Sprintf("%s >= %s", i.expr1.toString(), i.expr2.toString())
}
