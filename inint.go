package expr

import "fmt"

// InInt is the implementation of the operator 'InInt'
type InInt struct {
	*OperatorImpl
}

// NewInInt creates a new instance of InInt
func NewInInt() *InInt {
	inInt := &InInt{}
	inInt.OperatorImpl = NewOperator("IN", "EXPR.Operator.IN", booleanType, []interface{}{intType, "IN", intListType})
	return inInt
}

func (o *InInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *InInt) getName() string {
	return o.OperatorImpl.name
}

func (o *InInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *InInt) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'int in intlist' operator")
	return newInIntInstance(ctx)

}

type inIntInstance struct {
	*BooleanExpressionImpl
	eleExpr   IntExpression
	arrayExpr IntArrayExpression
}

func newInIntInstance(ctx OperatorContext) *inIntInstance {
	inInt := &inIntInstance{}

	inInt.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	inInt.eleExpr = ctx.getInputs()[0].(IntExpression)
	inInt.arrayExpr = ctx.getInputs()[1].(IntArrayExpression)

	return inInt
}

func (i *inIntInstance) evaluate(ctx EvaluationContext) bool {
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
func (i *inIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *inIntInstance) toString() string {
	return fmt.Sprintf("%s in %s", i.eleExpr.toString(), i.arrayExpr.toString())
}
