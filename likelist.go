package expr

import (
	"regexp"
	"strings"
)

// LikeList is the implementation of the operator 'LikeList'
type LikeList struct {
	*OperatorImpl
}

// NewLikeList creates a new instance of LikeList
func NewLikeList() *LikeList {
	likeList := &LikeList{}
	likeList.OperatorImpl = NewOperator("LikeList", "EXPR.Operator.LikeList", booleanType, []interface{}{stringListType, "LIKELIST", stringType})
	return likeList
}

func (o *LikeList) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *LikeList) getName() string {
	return o.OperatorImpl.name
}

func (o *LikeList) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *LikeList) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'strlist likelist str' operator")
	return newLikeListInstance(ctx)
}

type likeListInstance struct {
	*BooleanExpressionImpl
	expr      StringExpression
	arrayExpr StringArrayExpression
}

func newLikeListInstance(ctx OperatorContext) Expression {
	likeListInst := &likeListInstance{}

	likeListInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	likeListInst.arrayExpr = ctx.getInputs()[0].(StringArrayExpression)
	likeListInst.expr = ctx.getInputs()[1].(StringExpression)

	return likeListInst
}

func (i *likeListInstance) evaluate(ctx EvaluationContext) bool {
	str := i.expr.evaluate(ctx)
	if str == "" {
		return false
	}
	regExpPattern := strings.Replace(str, "%", "(.*)", -1)

	array := i.arrayExpr.evaluate(ctx)
	if len(array) > 0 {
		for i := 0; i < len(array); i++ {
			if matched, _ := regexp.MatchString(regExpPattern, array[i]); matched {
				return true
			}
		}
	}
	return false
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *likeListInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *likeListInstance) toString() string {
	return "like operator"
}
