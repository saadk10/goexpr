package expr

import "fmt"

// InStr is the implementation of the operator 'InStr'
type InStr struct {
	*OperatorImpl
}

// NewInStr creates a new instance of InStr
func NewInStr() *InStr {
	inStr := &InStr{}
	inStr.OperatorImpl = NewOperator("IN", "EXPR.Operator.IN", booleanType, []interface{}{stringType, "IN", stringListType})
	return inStr
}

func (o *InStr) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *InStr) getName() string {
	return o.OperatorImpl.name
}

func (o *InStr) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *InStr) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'str in str' operator")
	return newInStrInstance(ctx)

}

type inStrInstance struct {
	*BooleanExpressionImpl
	eleExpr   StringExpression
	arrayExpr StringArrayExpression
}

func newInStrInstance(ctx OperatorContext) Expression {
	inStr := &inStrInstance{}

	inStr.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)

	inStr.eleExpr = ctx.getInputs()[0].(StringExpression)
	inStr.arrayExpr = ctx.getInputs()[1].(StringArrayExpression)

	return inStr
}

func (i *inStrInstance) evaluate(ctx EvaluationContext) bool {
	str := i.eleExpr.evaluate(ctx)
	if str == "" {
		return false
	}
	array := i.arrayExpr.evaluate(ctx)
	for i := 0; i < len(array); i++ {
		member := array[i]
		if (member != "") && (str == member) {
			return true
		}
	}
	return false

}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *inStrInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *inStrInstance) toString() string {
	return fmt.Sprintf("%s in %s", i.eleExpr.toString(), i.arrayExpr.toString())
}
