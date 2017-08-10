package expr

import (
	"fmt"
)

// EqualsIntList is the implementation of the operator 'EqualsIntList'
type EqualsIntList struct {
	*OperatorImpl
}

// NewEqualsIntList creates a new instance of EqualsIntList
func NewEqualsIntList() *EqualsIntList {
	equalsStrList := &EqualsIntList{}
	equalsStrList.OperatorImpl = NewOperator("=", "EXPR.Operator.Equals", booleanType, []interface{}{intListType, "=", intType})
	return equalsStrList
}

func (o *EqualsIntList) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *EqualsIntList) getName() string {
	return o.OperatorImpl.name
}

func (o *EqualsIntList) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *EqualsIntList) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'intlist equals int' operator")
	return newEqualsIntListInstance(ctx)

}

type equalsIntListInstance struct {
	*BooleanExpressionImpl
	arrayExpr IntArrayExpression
	eleExpr   IntExpression
}

func newEqualsIntListInstance(ctx OperatorContext) Expression {
	equalsIntListInst := &equalsIntListInstance{}

	equalsIntListInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	equalsIntListInst.arrayExpr = ctx.getInputs()[0].(IntArrayExpression)
	equalsIntListInst.eleExpr = ctx.getInputs()[1].(IntExpression)

	return equalsIntListInst
}

func (i *equalsIntListInstance) evaluate(ctx EvaluationContext) bool {
	ele := i.eleExpr.evaluate(ctx)
	array := i.arrayExpr.evaluate(ctx)
	for i := 0; i < len(array); i++ {
		if ele == array[i] {
			return true
		}
	}
	return false
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *equalsIntListInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *equalsIntListInstance) toString() string {
	return fmt.Sprintf("%s = %s", i.arrayExpr.toString(), i.eleExpr.toString())
}
