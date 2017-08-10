package expr

import "fmt"

// LessThanEqualsInt is the implementation of the operator 'LessThanEqualsInt'
type LessThanEqualsInt struct {
	*OperatorImpl
}

// NewLessThanEqualsInt creates a new instance of LessThanEqualsInt
func NewLessThanEqualsInt() *LessThanEqualsInt {
	lte := &LessThanEqualsInt{}
	lte.OperatorImpl = NewOperator("<=", "EXPR.Operator.LessThanEquals", booleanType, []interface{}{intType, "<=", intType})
	return lte
}

func (o *LessThanEqualsInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *LessThanEqualsInt) getName() string {
	return o.OperatorImpl.name
}

func (o *LessThanEqualsInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *LessThanEqualsInt) createExpression(ctx OperatorContext) Expression {
	fmt.Println("Creating expression for 'int <= int' operator")
	return newLessThanEqualsIntInstance(ctx)

}

type lessThanEqualsIntInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 IntExpression
}

func newLessThanEqualsIntInstance(ctx OperatorContext) Expression {
	lteInst := &lessThanEqualsIntInstance{}

	lteInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	lteInst.expr1 = ctx.getInputs()[0].(IntExpression)
	lteInst.expr2 = ctx.getInputs()[1].(IntExpression)

	return lteInst
}

func (i *lessThanEqualsIntInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) <= i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *lessThanEqualsIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *lessThanEqualsIntInstance) toString() string {
	return fmt.Sprintf("%s <= %s", i.expr1.toString(), i.expr2.toString())
}
