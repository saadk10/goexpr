package expr

import (
	"fmt"
	"regexp"
	"strings"
)

// LikeIn is the implementation of the operator 'LikeIn'
type LikeIn struct {
	*OperatorImpl
}

// NewLikeIn creates a new instance of LikeIn
func NewLikeIn() *LikeIn {
	likeIn := &LikeIn{}
	likeIn.OperatorImpl = NewOperator("LikeIn", "EXPR.Operator.LikeIn", booleanType, []interface{}{stringType, "LIKEIN", stringListType})
	return likeIn
}

func (o *LikeIn) getInputElements() []interface{} {
	return o.OperatorImpl.getInputElements()
}

func (o *LikeIn) getName() string {
	return o.OperatorImpl.name
}

func (o *LikeIn) getReturnType() TypeImpl {
	return o.OperatorImpl.returnType
}

func (o *LikeIn) createExpression(ctx OperatorContext) Expression {
	log.Debugf("Creating expression for 'str likein strlist' operator")
	return newLikeInInstance(ctx)
}

type likeInInstance struct {
	*BooleanExpressionImpl
	expr      StringExpression
	arrayExpr StringArrayExpression
}

func newLikeInInstance(ctx OperatorContext) Expression {
	likeInInst := &likeInInstance{}

	likeInInst.BooleanExpressionImpl = NewBooleanExpression(ctx.ExpressionContext)
	likeInInst.expr = ctx.getInputs()[0].(StringExpression)
	likeInInst.arrayExpr = ctx.getInputs()[1].(StringArrayExpression)

	return likeInInst
}

func (i *likeInInstance) evaluate(ctx EvaluationContext) bool {
	str := i.expr.evaluate(ctx)
	if str == "" {
		return false
	}
	array := i.arrayExpr.evaluate(ctx)
	for i := 0; i < len(array); i++ {
		regExpPattern := strings.Replace(array[i], "%", "(.*)", -1)
		if matched, _ := regexp.MatchString(regExpPattern, str); matched {
			return true
		}
	}
	return false
}

func (i *likeInInstance) toString() string {
	return fmt.Sprintf("%s likein %s", i.expr.toString(), i.arrayExpr.toString())
}

// GetBoolean evaluates and returns the result for the boolean expression
func (i *likeInInstance) GetBoolean(ob interface{}) bool {
	ctx := EvaluationContextImpl{
		object: ob,
	}

	result := i.evaluate(&ctx)
	return result
}
