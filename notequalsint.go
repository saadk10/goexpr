package expr

import "fmt"

// NotEqualsInt is the implementation of the operator 'NotEqualsInt'
type NotEqualsInt struct {
	*OperatorImpl
}

// NewNotEqualsInt creates a new instance of NotEqualsInt
func NewNotEqualsInt() *NotEqualsInt {
	notEqualsInt := &NotEqualsInt{}
	notEqualsInt.OperatorImpl = NewOperator("!=", "EXPR.Operator.NotEquals", booleanType, []interface{}{intType, "!=", intType})
	return notEqualsInt
}

func (o *NotEqualsInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *NotEqualsInt) getName() string {
	return o.OperatorImpl.name
}

func (o *NotEqualsInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *NotEqualsInt) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'int != int' operator")
	return newNotEqualsIntInstance(ctx)

}

type notEqualsIntInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 IntExpression
}

func newNotEqualsIntInstance(ctx OperatorContext) Expression {
	notEqualsIntInst := &notEqualsIntInstance{}

	notEqualsIntInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	notEqualsIntInst.expr1 = ctx.getInputs()[0].(IntExpression)
	notEqualsIntInst.expr2 = ctx.getInputs()[1].(IntExpression)

	return notEqualsIntInst
}

func (i *notEqualsIntInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) != i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *notEqualsIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *notEqualsIntInstance) toString() string {
	return fmt.Sprintf("%s != %s", i.expr1.toString(), i.expr2.toString())
}
