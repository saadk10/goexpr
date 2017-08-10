package expr

import "fmt"

// Not is the implementation of the operator 'Not'
type Not struct {
	*OperatorImpl
}

// NewNot creates a new instance of Not
func NewNot() *Not {
	not := &Not{}
	not.OperatorImpl = NewOperator("Not", "EXPR.Operator.Not", booleanType, []interface{}{"NOT", booleanType})
	return not
}

func (o *Not) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *Not) getName() string {
	return o.OperatorImpl.name
}

func (o *Not) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *Not) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'not bool' operator")
	return newNotInstance(ctx)

}

type notInstance struct {
	*BooleanExpressionImpl
	expr1 BooleanExpression
}

func newNotInstance(ctx OperatorContext) Expression {
	notInst := &notInstance{}

	notInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	notInst.expr1 = ctx.getInputs()[0].(BooleanExpression)

	return notInst
}

func (i *notInstance) evaluate(ctx EvaluationContext) bool {
	result := i.expr1.evaluate(ctx)
	return !result
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *notInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *notInstance) toString() string {
	return fmt.Sprintf("NOT %s", i.expr1.toString())
}
