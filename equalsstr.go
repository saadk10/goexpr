package expr

import (
	"fmt"
)

// EqualsStr is the implementation of the operator 'EqualsStr'
type EqualsStr struct {
	*OperatorImpl
}

// NewEqualsStr creates a new instance of EqualsStr
func NewEqualsStr() *EqualsStr {
	equalsStr := &EqualsStr{}
	equalsStr.OperatorImpl = NewOperator("=", "EXPR.Operator.Equals", booleanType, []interface{}{stringType, "=", stringType})
	return equalsStr
}

func (o *EqualsStr) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *EqualsStr) getName() string {
	return o.OperatorImpl.name
}

func (o *EqualsStr) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *EqualsStr) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'str equals str' operator")
	return newEqualsStrInstance(ctx)
}

func (o *EqualsStr) toString() string {
	return o.name
}

type equalsStrInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 StringExpression
}

func newEqualsStrInstance(ctx OperatorContext) Expression {
	equalsStrInst := &equalsStrInstance{}

	equalsStrInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	equalsStrInst.expr1 = ctx.getInputs()[0].(StringExpression)
	equalsStrInst.expr2 = ctx.getInputs()[1].(StringExpression)

	return equalsStrInst
}

func (i *equalsStrInstance) evaluate(ctx EvaluationContext) bool {
	str1 := i.expr1.evaluate(ctx)
	if str1 == any {
		return true
	}
	str2 := i.expr2.evaluate(ctx)
	if str2 == any {
		return true
	}
	if str1 == str2 {
		return true
	}
	if str1 == "" || str2 == "" {
		return false
	}
	return false
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *equalsStrInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)

	return result
}

func (i *equalsStrInstance) toString() string {
	return fmt.Sprintf("%s = %s", i.expr1.toString(), i.expr2.toString())
}
