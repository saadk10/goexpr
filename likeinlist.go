package expr

import (
	"fmt"
	"regexp"
	"strings"
)

// LikeInList is the implementation of the operator 'LikeIn'
type LikeInList struct {
	*OperatorImpl
}

// NewLikeInList creates a new instance of LikeIn
func NewLikeInList() *LikeInList {
	likeInList := &LikeInList{}
	likeInList.OperatorImpl = NewOperator("LikeIn", "EXPR.Operator.LikeIn", booleanType, []interface{}{stringListType, "LIKEIN", stringListType})
	return likeInList
}

func (o *LikeInList) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *LikeInList) getName() string {
	return o.OperatorImpl.name
}

func (o *LikeInList) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *LikeInList) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'strlist likein strlist' operator")
	return newLikeInListInstance(ctx)
}

type likeInListInstance struct {
	*BooleanExpressionImpl
	expr      StringArrayExpression
	arrayExpr StringArrayExpression
}

func newLikeInListInstance(ctx OperatorContext) Expression {
	likeInInstList := &likeInListInstance{}

	likeInInstList.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	likeInInstList.expr = ctx.getInputs()[0].(StringArrayExpression)
	likeInInstList.arrayExpr = ctx.getInputs()[1].(StringArrayExpression)

	return likeInInstList
}

func (i *likeInListInstance) evaluate(ctx EvaluationContext) bool {
	strs := i.expr.evaluate(ctx)
	if len(strs) == 0 {
		return false
	}
	array := i.arrayExpr.evaluate(ctx)
	for j := 0; j < len(strs); j++ {
		for i := 0; i < len(array); i++ {
			regExpPattern := strings.Replace(array[i], "%", "(.*)", -1)
			if matched, _ := regexp.MatchString(regExpPattern, strs[j]); matched {
				return true
			}
		}
	}
	return false
}

func (i *likeInListInstance) toString() string {
	return fmt.Sprintf("%s likein %s", i.expr.toString(), i.arrayExpr.toString())
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *likeInListInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}
