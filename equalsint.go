package expr

import "fmt"

// EqualsInt is the implementation of the operator 'EqualsInt'
type EqualsInt struct {
	*OperatorImpl
}

// NewEqualsInt creates a new instance of EqualsInt
func NewEqualsInt() *EqualsInt {
	equalsInt := &EqualsInt{}
	equalsInt.OperatorImpl = NewOperator("=", "EXPR.Operator.Equals", booleanType, []interface{}{intType, "=", intType})
	return equalsInt
}

func (o *EqualsInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *EqualsInt) getName() string {
	return o.OperatorImpl.name
}

func (o *EqualsInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *EqualsInt) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'int equals int' operator")
	return newEqualsIntInstance(ctx)

}

type equalsIntInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 IntExpression
}

func newEqualsIntInstance(ctx OperatorContext) Expression {
	equalsStrInstance := &equalsIntInstance{}

	equalsStrInstance.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	equalsStrInstance.expr1 = ctx.getInputs()[0].(IntExpression)
	equalsStrInstance.expr2 = ctx.getInputs()[1].(IntExpression)

	return equalsStrInstance
}

func (i *equalsIntInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) == i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *equalsIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *equalsIntInstance) toString() string {
	return fmt.Sprintf("%s = %s", i.expr1.toString(), i.expr2.toString())
}
