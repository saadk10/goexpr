package expr

import (
	"fmt"
	"regexp"
	"strings"
)

// Like is the implementation of the operator 'Like'
type Like struct {
	*OperatorImpl
}

// NewLike creates a new instance of Like
func NewLike() *Like {
	like := &Like{}
	like.OperatorImpl = NewOperator("Like", "EXPR.Operator.Like", booleanType, []interface{}{stringType, "LIKE", stringType})
	return like
}

func (o *Like) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *Like) getName() string {
	return o.OperatorImpl.name
}

func (o *Like) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *Like) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'str like str' operator")
	return newLikeInstance(ctx)
}

type likeInstance struct {
	*BooleanExpressionImpl
	expr1, expr2 StringExpression
}

func newLikeInstance(ctx OperatorContext) Expression {
	likeInst := &likeInstance{}

	likeInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	likeInst.expr1 = ctx.getInputs()[0].(StringExpression)
	likeInst.expr2 = ctx.getInputs()[1].(StringExpression)

	return likeInst
}

func (i *likeInstance) evaluate(ctx EvaluationContext) bool {
	pattern := i.expr2.evaluate(ctx)
	regExpPattern := strings.Replace(pattern, "%", "(.*)", -1)

	str := i.expr1.evaluate(ctx)
	if str == "" {
		return false
	}
	result, _ := regexp.MatchString(regExpPattern, str)

	return result
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *likeInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}

func (i *likeInstance) toString() string {
	return fmt.Sprintf("%s like %s", i.expr1.toString(), i.expr2.toString())
}
