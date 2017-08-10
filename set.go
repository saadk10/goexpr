package expr

import "fmt"

// Set is the implementation of the operator 'Set'
type Set struct {
	*OperatorImpl
}

// NewSet creates a new instance of Set
func NewSet() *Set {
	set := &Set{}
	set.OperatorImpl = NewOperator("Set", "EXPR.Operator.Set", booleanType, []interface{}{"SET", stringType})
	return set
}

func (a *Set) getInputElements() []interface{} {
	return a.OperatorImpl.getInputElements()
}

func (a *Set) getName() string {
	return a.OperatorImpl.name
}

func (a *Set) getReturnType() TypeImpl {
	return a.OperatorImpl.returnType
}

func (a *Set) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'set str' operator")
	return newSetInstance(ctx)

}

type setInstance struct {
	*BooleanExpressionImpl
	expr1 StringExpression
}

func newSetInstance(ctx OperatorContext) Expression {
	setInst := &setInstance{}

	setInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	setInst.expr1 = ctx.getInputs()[0].(StringExpression)

	return setInst
}

func (i *setInstance) evaluate(ctx EvaluationContext) bool {
	return i.expr1.evaluate(ctx) != ""
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *setInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *setInstance) toString() string {
	return fmt.Sprintf("set %s", i.expr1.toString())
}
