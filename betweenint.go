package expr

import (
	"fmt"
)

// BetweenInt is the implementation of the operator 'BetweenInt'
type BetweenInt struct {
	*OperatorImpl
}

// NewBetweenInt creates a new instance of BetweenInt
func NewBetweenInt() *BetweenInt {
	betweenInt := &BetweenInt{}
	betweenInt.OperatorImpl = NewOperator("Between", "EXPR.Operator.Between", booleanType, []interface{}{intType, "BETWEEN", intType, "AND", intType})
	return betweenInt
}

func (o *BetweenInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *BetweenInt) getName() string {
	return o.OperatorImpl.name
}

func (o *BetweenInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *BetweenInt) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'int between int and int' operator")
	return newBetweenIntInstance(ctx)

}

type betweenIntInstance struct {
	*BooleanExpressionImpl
	expr, expr1, expr2 IntExpression
}

func newBetweenIntInstance(ctx OperatorContext) Expression {
	betweenIntInst := &betweenIntInstance{}

	betweenIntInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	betweenIntInst.expr = ctx.getInputs()[0].(IntExpression)
	betweenIntInst.expr1 = ctx.getInputs()[1].(IntExpression)
	betweenIntInst.expr2 = ctx.getInputs()[2].(IntExpression)

	return betweenIntInst
}

func (i *betweenIntInstance) evaluate(ctx EvaluationContext) bool {
	val := i.expr.evaluate(ctx)
	val1 := i.expr1.evaluate(ctx)
	if val < val1 {
		return false
	}
	val2 := i.expr2.evaluate(ctx)
	return val <= val2
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *betweenIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *betweenIntInstance) toString() string {
	return fmt.Sprintf("between %s AND %s", i.expr1.toString(), i.expr2.toString())
}
