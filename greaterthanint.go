package expr

import "fmt"

// GreaterThanInt is the implementation of the operator 'GreaterThanInt'
type GreaterThanInt struct {
	*OperatorImpl
}

// NewGreaterThanInt creates a new instance of GreaterThanInt
func NewGreaterThanInt() *GreaterThanInt {
	gt := &GreaterThanInt{}
	gt.OperatorImpl = NewOperator(">", "EXPR.Operator.GreaterThan", booleanType, []interface{}{intType, ">", intType})
	return gt
}

func (o *GreaterThanInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *GreaterThanInt) getName() string {
	return o.OperatorImpl.name
}

func (o *GreaterThanInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *GreaterThanInt) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'int > int' operator")
	return newGreaterThanIntInstance(ctx)

}

type greaterThanIntInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 IntExpression
}

func newGreaterThanIntInstance(ctx OperatorContext) Expression {
	gtInst := &greaterThanIntInstance{}

	gtInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	gtInst.expr1 = ctx.getInputs()[0].(IntExpression)
	gtInst.expr2 = ctx.getInputs()[1].(IntExpression)

	return gtInst
}

func (i *greaterThanIntInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) > i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *greaterThanIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *greaterThanIntInstance) toString() string {
	return fmt.Sprintf("%s > %s", i.expr1.toString(), i.expr2.toString())
}
