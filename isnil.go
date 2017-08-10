package expr

import "fmt"

// IsNil is the implementation of the operator 'IsNil'
type IsNil struct {
	*OperatorImpl
}

// NewIsNil creates a new instance of IsNil
func NewIsNil() *IsNil {
	isNil := &IsNil{}
	isNil.OperatorImpl = NewOperator("ISNIL", "EXPR.Operator.ISNIL", booleanType, []interface{}{stringType, "ISNIL"})
	return isNil
}

func (o *IsNil) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *IsNil) getName() string {
	return o.OperatorImpl.name
}

func (o *IsNil) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *IsNil) createExpression(ctx OperatorContext) Expression {
	fmt.Println("Creating expression for 'str isnil' operator")
	return newIsNilInstance(ctx)

}

type isNilInstance struct {
	*BooleanExpressionImpl
	expr StringExpression
}

func newIsNilInstance(ctx OperatorContext) Expression {
	isNilInst := &isNilInstance{}

	isNilInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	isNilInst.expr = ctx.getInputs()[0].(StringExpression)

	return isNilInst
}

func (i *isNilInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr.evaluate(ctx) == ""
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *isNilInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *isNilInstance) toString() string {
	return fmt.Sprintf("%s isnil", i.expr.toString())
}
