package expr

import "fmt"

// NotEqualsStr is the implementation of the operator 'NotEqualsStr'
type NotEqualsStr struct {
	*OperatorImpl
}

// NewNotEqualsStr creates a new instance of NotEqualsStr
func NewNotEqualsStr() *NotEqualsStr {
	notEqualsStr := &NotEqualsStr{}
	notEqualsStr.OperatorImpl = NewOperator("!=", "EXPR.Operator.NotEquals", booleanType, []interface{}{stringType, "!=", stringType})
	return notEqualsStr
}

func (o *NotEqualsStr) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *NotEqualsStr) getName() string {
	return o.OperatorImpl.name
}

func (o *NotEqualsStr) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *NotEqualsStr) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'str != str' operator")
	return newNotEqualsStrInstance(ctx)

}

type notEqualsStrInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 StringExpression
}

func newNotEqualsStrInstance(ctx OperatorContext) Expression {
	notEqualsStrInst := &notEqualsStrInstance{}

	notEqualsStrInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	notEqualsStrInst.expr1 = ctx.getInputs()[0].(StringExpression)
	notEqualsStrInst.expr2 = ctx.getInputs()[1].(StringExpression)

	return notEqualsStrInst
}

func (i *notEqualsStrInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) != i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *notEqualsStrInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *notEqualsStrInstance) toString() string {
	return fmt.Sprintf("%s != %s", i.expr1.toString(), i.expr2.toString())
}
