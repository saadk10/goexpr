package expr

import (
	"fmt"
	"regexp"
	"strings"
)

// LikeIgnoreCase is the implementation of the operator 'LikeIgnoreCase'
type LikeIgnoreCase struct {
	*OperatorImpl
}

// NewLikeIgnoreCase creates a new instance of LikeIgnoreCase
func NewLikeIgnoreCase() *LikeIgnoreCase {
	like := &LikeIgnoreCase{}
	like.OperatorImpl = NewOperator("LikeIgnoreCase", "EXPR.Operator.LikeIgnoreCase", booleanType, []interface{}{stringType, "LIKEIGNORECASE", stringType})
	return like
}

func (o *LikeIgnoreCase) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *LikeIgnoreCase) getName() string {
	return o.OperatorImpl.name
}

func (o *LikeIgnoreCase) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *LikeIgnoreCase) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'str likeignorecase str' operator")
	return newLikeIgnoreCaseInstance(ctx)
}

type likeIgnoreCaseInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 StringExpression
}

func newLikeIgnoreCaseInstance(ctx OperatorContext) Expression {
	likeIgnoreCaseInstance := &likeIgnoreCaseInstance{}

	likeIgnoreCaseInstance.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	likeIgnoreCaseInstance.expr1 = ctx.getInputs()[0].(StringExpression)
	likeIgnoreCaseInstance.expr2 = ctx.getInputs()[1].(StringExpression)

	return likeIgnoreCaseInstance
}

func (i *likeIgnoreCaseInstance) evaluate(ctx EvaluationContext) bool {
	pattern := strings.ToLower(i.expr2.evaluate(ctx))
	regExpPattern := strings.Replace(pattern, "%", "(.*)", -1)

	str := strings.ToLower(i.expr1.evaluate(ctx))
	if str == "" {
		return false
	}
	result, _ := regexp.MatchString(regExpPattern, str)

	return result
}

func (i *likeIgnoreCaseInstance) toString() string {
	return fmt.Sprintf("%s likeignorecase %s", i.expr1.toString(), i.expr2.toString())
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *likeIgnoreCaseInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}
