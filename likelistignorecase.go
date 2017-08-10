package expr

import (
	"regexp"
	"strings"
)

// LikeListIgnoreCase is the implementation of the operator 'LikeList'
type LikeListIgnoreCase struct {
	*OperatorImpl
}

// NewLikeListIgnoreCase creates a new instance of LikeList
func NewLikeListIgnoreCase() *LikeListIgnoreCase {
	likeListIgnoreCase := &LikeListIgnoreCase{}
	likeListIgnoreCase.OperatorImpl = NewOperator("LikeList", "EXPR.Operator.LikeList", booleanType, []interface{}{stringListType, "LIKELIST", stringType})
	return likeListIgnoreCase
}

func (o *LikeListIgnoreCase) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *LikeListIgnoreCase) getName() string {
	return o.OperatorImpl.name
}

func (o *LikeListIgnoreCase) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *LikeListIgnoreCase) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'strlist likelistignorecase str' operator")
	return newLikeListInstance(ctx)
}

type likeListIgnoreCaseInstance struct {
	*BooleanExpressionImpl
	expr      StringExpression
	arrayExpr StringArrayExpression
}

func newLikeListIgnoreCaseInstance(ctx OperatorContext) Expression {
	likeListInst := &likeListIgnoreCaseInstance{}

	likeListInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	likeListInst.arrayExpr = ctx.getInputs()[0].(StringArrayExpression)
	likeListInst.expr = ctx.getInputs()[1].(StringExpression)

	return likeListInst
}

func (i *likeListIgnoreCaseInstance) evaluate(ctx EvaluationContext) bool {
	str := i.expr.evaluate(ctx)
	if str == "" {
		return false
	}
	regExpPattern := strings.Replace(str, "%", "(.*)", -1)

	array := i.arrayExpr.evaluate(ctx)
	if len(array) > 0 {
		for i := 0; i < len(array); i++ {
			if matched, _ := regexp.MatchString(strings.ToLower(regExpPattern), strings.ToLower(array[i])); matched {
				return true
			}
		}
	}
	return false
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *likeListIgnoreCaseInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *likeListIgnoreCaseInstance) toString() string {
	return "like operator"
}
