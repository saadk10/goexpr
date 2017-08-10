package expr

import "fmt"

// LessThanInt is the implementation of the operator 'LessThanInt'
type LessThanInt struct {
	*OperatorImpl
}

// NewLessThanInt creates a new instance of LessThanInt
func NewLessThanInt() *LessThanInt {
	lt := &LessThanInt{}
	lt.OperatorImpl = NewOperator("<", "EXPR.Operator.<", booleanType, []interface{}{intType, "<", intType})
	return lt
}

func (o *LessThanInt) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *LessThanInt) getName() string {
	return o.OperatorImpl.name
}

func (o *LessThanInt) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *LessThanInt) createExpression(ctx OperatorContext) Expression {
	fmt.Println("Creating expression for 'int < int' operator")
	return newLessThanIntInstance(ctx)

}

type lessThanIntInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 IntExpression
}

func newLessThanIntInstance(ctx OperatorContext) Expression {
	ltInst := &lessThanIntInstance{}

	ltInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	ltInst.expr1 = ctx.getInputs()[0].(IntExpression)
	ltInst.expr2 = ctx.getInputs()[1].(IntExpression)

	return ltInst
}

func (i *lessThanIntInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) < i.expr2.evaluate(ctx)
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *lessThanIntInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *lessThanIntInstance) toString() string {
	return fmt.Sprintf("%s < %s", i.expr1.toString(), i.expr2.toString())
}
