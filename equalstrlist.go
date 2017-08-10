package expr

import (
	"fmt"
)

// EqualsStrList is the implementation of the operator 'EqualsStrList'
type EqualsStrList struct {
	*OperatorImpl
}

// NewEqualsStrList creates a new instance of EqualsStrList
func NewEqualsStrList() *EqualsStrList {
	equalsStrList := &EqualsStrList{}
	equalsStrList.OperatorImpl = NewOperator("=", "EXPR.Operator.Equals", booleanType, []interface{}{stringListType, "=", stringType})
	return equalsStrList
}

func (o *EqualsStrList) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *EqualsStrList) getName() string {
	return o.OperatorImpl.name
}

func (o *EqualsStrList) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *EqualsStrList) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for equal str list")
	return newEqualsStrListInstance(ctx)

}

type equalsStrListInstance struct {
	*BooleanExpressionImpl
	expr1 StringArrayExpression
	expr2 StringExpression
}

func newEqualsStrListInstance(ctx OperatorContext) Expression {
	equalsStrListInst := &equalsStrListInstance{}

	equalsStrListInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	equalsStrListInst.expr1 = ctx.getInputs()[0].(StringArrayExpression)
	equalsStrListInst.expr2 = ctx.getInputs()[1].(StringExpression)

	return equalsStrListInst
}

func (i *equalsStrListInstance) evaluate(ctx EvaluationContext) bool {
	str1 := i.expr1.evaluate(ctx)
	str2 := i.expr2.evaluate(ctx)
	str2empty := (str2 == "")
	if len(str1) == 0 {
		if str2empty {
			return true
		}
		return false
	}
	for i := 0; i < len(str1); i++ {
		s := str1[i]
		if s == "" {
			if str2empty {
				return true
			}
		} else if !str2empty {
			if s != "" && s == str2 {
				return true
			}
		}
	}
	return false
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *equalsStrListInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)

	return result
}

func (i *equalsStrListInstance) toString() string {
	return fmt.Sprintf("%s = %s", i.expr1.toString(), i.expr2.toString())
}
