package expr

import "fmt"

// IsNotNil is the implementation of the operator 'IsNotNil'
type IsNotNil struct {
	*OperatorImpl
}

// NewIsNotNil creates a new instance of IsNotNil
func NewIsNotNil() *IsNotNil {
	isNotNil := &IsNotNil{}
	isNotNil.OperatorImpl = NewOperator("ISNOTNIL", "EXPR.Operator.ISNOTNIL", booleanType, []interface{}{stringType, "ISNOTNIL"})
	return isNotNil
}

func (o *IsNotNil) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *IsNotNil) getName() string {
	return o.OperatorImpl.name
}

func (o *IsNotNil) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *IsNotNil) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'str isnotnil' operator")
	return newIsNotNilInstance(ctx)

}

type isNotNilInstance struct {
	*BooleanExpressionImpl
	expr StringExpression
}

func newIsNotNilInstance(ctx OperatorContext) Expression {
	isNotNilInst := &isNotNilInstance{}

	isNotNilInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	isNotNilInst.expr = ctx.getInputs()[0].(StringExpression)

	return isNotNilInst
}

func (i *isNotNilInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr.evaluate(ctx) != ""
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *isNotNilInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *isNotNilInstance) toString() string {
	return fmt.Sprintf("%s isnotnil", i.expr.toString())
}
