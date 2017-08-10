package expr

import (
	"fmt"
	"strconv"
	"testing"
)

func TestExpr(t *testing.T) {
	lm := &LanguageManagerImpl{}
	lang, _ := lm.CreateLanguage("test", false)
	testOperand := NewTestOperand()
	testStringOperand := NewTestStringOperand()
	testIntOperand := NewTestIntOperand()
	testStringArrayOperand := NewTestStringArrayOperand()

	lang.RegisterOperand(testOperand)
	lang.RegisterOperand(testStringOperand)
	lang.RegisterOperand(testIntOperand)
	lang.RegisterOperand(testStringArrayOperand)

	var expr BooleanExpression

	expr, _ = lang.CreateBooleanExpression("(test = 'test')")
	check(expr.GetBoolean("test"), true, "str equals true", t)
	check(expr.GetBoolean("notest"), false, "str equals true", t)

	expr, _ = lang.CreateBooleanExpression("(test = 10)")
	check(expr.GetBoolean("10"), true, "int equals true", t)
	check(expr.GetBoolean("11"), false, "int equals true", t)

	expr, _ = lang.CreateBooleanExpression("(NOT (test = 'ab'))")
	check(expr.GetBoolean("foo"), true, "str not equal true", t)
	check(expr.GetBoolean("ab"), false, "str not equal false", t)

	expr, _ = lang.CreateBooleanExpression("(NOT (test = 'ab')) and (test = 'test')")
	check(expr.GetBoolean("foo"), false, "str not equal true", t)
	check(expr.GetBoolean("test"), true, "str not equal false", t)

	expr, _ = lang.CreateBooleanExpression("(test IN ('foo', 'bar', 'xyz'))")
	check(expr.GetBoolean("foo"), true, "str1 in list true", t)
	check(expr.GetBoolean("ab"), false, "str1 in list true", t)
	check(expr.GetBoolean("xyz"), true, "str1 in list true", t)

	expr, _ = lang.CreateBooleanExpression("(test = 'ab') OR (test = 'test')")
	check(expr.GetBoolean("foo"), false, "str or false", t)
	check(expr.GetBoolean("test"), true, "str or true", t)

	expr, _ = lang.CreateBooleanExpression("(('a','b','c') = test)")
	check(expr.GetBoolean("a"), true, "str or false", t)
	check(expr.GetBoolean("d"), false, "str or false", t)

	expr, _ = lang.CreateBooleanExpression("test between 5 and 10")
	check(expr.GetBoolean("1"), false, "between str less", t)
	check(expr.GetBoolean("5"), true, "between str bottom", t)
	check(expr.GetBoolean("7"), true, "between str mid", t)
	check(expr.GetBoolean("10"), true, "between str top", t)
	check(expr.GetBoolean("11"), false, "between str more", t)

	expr, _ = lang.CreateBooleanExpression("((1,2,3) = test)")
	check(expr.GetBoolean("1"), true, "str list equals false", t)
	check(expr.GetBoolean("3"), true, "str list equals true", t)
	check(expr.GetBoolean("4"), false, "str list equals false", t)

	expr, _ = lang.CreateBooleanExpression("true")
	check(expr.GetBoolean(nil), true, "constant true", t)

	expr, _ = lang.CreateBooleanExpression("false")
	check(expr.GetBoolean(nil), false, "constant false", t)

	expr, _ = lang.CreateBooleanExpression("(test LIKE '%a%b')")
	check(expr.GetBoolean("a123b"), true, "str like true", t)
	check(expr.GetBoolean("foo"), false, "str like false", t)
	check(expr.GetBoolean("A123b"), false, "str like false", t)

	expr, _ = lang.CreateBooleanExpression("(test LIKEIGNORECASE '%a%b')")
	check(expr.GetBoolean("a123b"), true, "str like true", t)
	check(expr.GetBoolean("foo"), false, "str like false", t)
	check(expr.GetBoolean("A123b"), true, "str like false", t)

	expr, _ = lang.CreateBooleanExpression("(test LIKEIN ('a%z', 'A%Z'))")
	check(expr.GetBoolean("ABCDEFZ"), true, "str like list true", t)
	check(expr.GetBoolean("ABCDEFG"), false, "str like list false", t)

	expr, _ = lang.CreateBooleanExpression("(test >= 10)")
	check(expr.GetBoolean("11"), true, "int greater than equal true", t)
	check(expr.GetBoolean("10"), true, "str greater than equal false", t)
	check(expr.GetBoolean("9"), false, "str greater than equal false", t)

	expr, _ = lang.CreateBooleanExpression("(test > 10)")
	check(expr.GetBoolean("12"), true, "int greater than equal true", t)
	check(expr.GetBoolean("10"), false, "str greater than equal false", t)

	expr, _ = lang.CreateBooleanExpression("(test <= 10)")
	check(expr.GetBoolean("9"), true, "str less than equal false", t)
	check(expr.GetBoolean("10"), true, "str less than equal false", t)
	check(expr.GetBoolean("11"), false, "int less than equal true", t)

	expr, _ = lang.CreateBooleanExpression("(test < 10)")
	check(expr.GetBoolean("8"), true, "int less than equal true", t)
	check(expr.GetBoolean("10"), false, "str less than equal false", t)

	expr, _ = lang.CreateBooleanExpression("test isnotnil")
	check(expr.GetBoolean("abc"), true, "str is nil false", t)
	check(expr.GetBoolean(""), false, "int in list false", t)

	expr, _ = lang.CreateBooleanExpression("test isnil")
	check(expr.GetBoolean(""), true, "str is nil false", t)
	check(expr.GetBoolean("abc"), false, "int in list false", t)

	expr, _ = lang.CreateBooleanExpression("test != 10")
	check(expr.GetBoolean("11"), true, "int is not equals true", t)
	check(expr.GetBoolean("10"), false, "int in not equals false", t)

	expr, _ = lang.CreateBooleanExpression("test != 'test'")
	check(expr.GetBoolean("nottest"), true, "str is not equals true", t)
	check(expr.GetBoolean("test"), false, "str in not equals false", t)

	expr, _ = lang.CreateBooleanExpression("(SET test)")
	check(expr.GetBoolean("set"), true, "str true", t)
	check(expr.GetBoolean(""), false, "str false", t)

	expr, _ = lang.CreateBooleanExpression("(test IN (1,2,3))")
	check(expr.GetBoolean("2"), true, "int in list true", t)
	check(expr.GetBoolean("4"), false, "int in list false", t)

	expr, _ = lang.CreateBooleanExpression("(testString like '%.ibm.com')")
	check(expr.GetBoolean("foo.ibm.com"), true, "str pattern true", t)
	check(expr.GetBoolean("foo.acme.com"), false, "str pattern false", t)

	expr, _ = lang.CreateBooleanExpression("testint between 5 and 10")
	check(expr.GetBoolean(1), false, "between str less", t)
	check(expr.GetBoolean(5), true, "between str bottom", t)
	check(expr.GetBoolean(7), true, "between str mid", t)
	check(expr.GetBoolean(10), true, "between str top", t)
	check(expr.GetBoolean(11), false, "between str more", t)

	expr, _ = lang.CreateBooleanExpression("(testStringArray = testString)")
	check(expr.GetBoolean("THREE"), true, "str[] = true", t)
	check(expr.GetBoolean("SIX"), false, "str[] = false", t)

	expr, _ = lang.CreateBooleanExpression("(('abc', 'xyz') LIKELIST test)")
	check(expr.GetBoolean("a%c"), true, "str list like true", t)
	check(expr.GetBoolean("a%d"), false, "str list like false", t)
	check(expr.GetBoolean("A%d"), false, "str list like false", t)

	expr, _ = lang.CreateBooleanExpression("(testStringArray LIKEIN ('a%b', '%IVE'))")
	check(expr.GetBoolean(""), true, "strlist likein strlist true", t)

	expr, _ = lang.CreateBooleanExpression("(testStringArray LIKEIN ('a%b', 'A%B'))")
	check(expr.GetBoolean(""), false, "strlist likein strlist false", t)
}

func check(val bool, expectedVal bool, testName string, t *testing.T) {
	if val != expectedVal {
		t.Fatalf("Test '%s' failed", testName)
	}
	fmt.Println("Test Passed")
}

type TestStringOperand struct {
	*OperandImpl
}

func NewTestStringOperand() *TestStringOperand {
	oper := &TestStringOperand{}
	oper.OperandImpl = NewOperandWithNameType("testString", stringType)

	return oper
}

func (o *TestStringOperand) CreateExpression(ctx OperandContext) Expression {
	return NewString(ctx)
}

// Test operand of integer type
type TestIntOperand struct {
	*OperandImpl
}

func NewTestIntOperand() *TestIntOperand {
	oper := &TestIntOperand{}

	oper.OperandImpl = NewOperandWithNameType("testint", intType)
	return oper
}

func (o *TestIntOperand) CreateExpression(ctx OperandContext) Expression {
	return NewInt(ctx)
}

// Test operand of string array type
type TestStringArrayOperand struct {
	*OperandImpl
}

func NewTestStringArrayOperand() *TestStringArrayOperand {
	oper := &TestStringArrayOperand{}

	typ := stringListType.TypeImpl
	oper.OperandImpl = NewOperandWithNameType("testStringArray", typ)
	return oper
}

func (o *TestStringArrayOperand) CreateExpression(ctx OperandContext) Expression {
	return NewMyStringArray(ctx)
}

type MyStringArray struct {
	StringArrayExpressionImpl
	vals []string
}

func NewMyStringArray(ctx OperandContext) *MyStringArray {
	myStringArray := &MyStringArray{}

	myStringArray.StringArrayExpressionImpl.ExpressionImpl = *NewExpression(ctx.ExpressionContext)
	myStringArray.vals = []string{"ONE", "TWO", "THREE", "FOUR", "FIVE"}
	return myStringArray
}

func (s *MyStringArray) evaluate(ctx EvaluationContext) []string {
	return s.vals
}

func (s *MyStringArray) toString() string {
	return s.StringArrayExpressionImpl.ExpressionImpl.Str
}

// Test operand of both string and integer type
type TestOperand struct {
	*OperandImpl
}

func NewTestOperand() *TestOperand {
	oper := &TestOperand{}
	oper.OperandImpl = NewOperandWithNameTypeList("test", []TypeImpl{stringType, intType})

	return oper
}

func (o *TestOperand) CreateExpression(ctx OperandContext) Expression {
	if ctx.GetType() == stringType {
		return NewString(ctx)
	} else if ctx.GetType() == intType {
		return NewInt(ctx)
	}
	return nil
}

func (o TestOperand) toString() string {
	return o.OperandImpl.toString()
}

type MyString struct {
	StringExpressionImpl
}

func NewString(ctx OperandContext) *MyString {
	myString := &MyString{}

	myString.StringExpressionImpl = *NewStringExpression(ctx.ExpressionContext)
	return myString

}

func (s *MyString) evaluate(ctx EvaluationContext) string {
	return ctx.GetObject().(string)
}

func (s *MyString) toString() string {
	return s.StringExpressionImpl.ExpressionImpl.Str
}

type MyInt struct {
	IntExpressionImpl
}

func NewInt(ctx OperandContext) *MyInt {
	myInt := &MyInt{}

	myInt.IntExpressionImpl = *NewIntExpression(ctx.ExpressionContext)
	return myInt
}

func (i *MyInt) evaluate(ctx EvaluationContext) int {
	str, ok := ctx.GetObject().(string)
	var val int
	if ok {
		val, _ = strconv.Atoi(str)
	} else {
		val = ctx.GetObject().(int)
	}

	return val
}

func (i *MyInt) toString() string {
	return i.IntExpressionImpl.ExpressionImpl.Str
}
