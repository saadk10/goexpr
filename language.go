package expr

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("expr")
)

// Language defines the interface
type Language interface {
	RegisterOperator(operator Operator) error
	RegisterOperand(operand Operand) error
	CreateBooleanExpression(exprString string) (BooleanExpression, error)
}

// LanguageImpl ..
type LanguageImpl struct {
	initialized             bool
	defaultOperators        []interface{}
	operatorList            []interface{}
	defaultOperands         []interface{}
	operandList             []interface{}
	operandMap              map[string]interface{}
	operatorMap             map[string]interface{}
	registerDefaultOperands bool
	operands                []Operand
	operators               []Operator
	name                    string
	operatorIdentifiers     map[string]string
	fieldSeparator          string
	ctxList                 []interface{}
	debug                   bool
}

// NewLanguage returns an instance of the implementation of Language
func NewLanguage() (*LanguageImpl, error) {
	l := &LanguageImpl{}
	err := l.init()
	if err != nil {
		return nil, fmt.Errorf("Error occured during initializing of new language: %s", err)
	}

	return l, nil
}

// Initialize language struct
func (l *LanguageImpl) init() error {
	if l.debug {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.INFO, "")
	}

	log.Infof("Initializing language '%s'\n", l.name)
	if l.initialized {
		return nil
	}

	l.defaultOperators = make([]interface{}, 0)
	l.defaultOperators = append(l.defaultOperators, NewEqualsStrList(), NewEqualsStr(), NewTrue(), NewEqualsInt(), NewNot(), NewAnd(), NewInStr(), NewEqualsStrList(), NewOr(), NewBetweenInt(), NewEqualsIntList(), NewFalse(), NewLike(), NewLikeIn(), NewGreaterThanEqualsInt(), NewGreaterThanInt(), NewLessThanEqualsInt(), NewLessThanInt(), NewInInt(), NewIsNotNil(), NewIsNil(), NewNotEqualsInt(), NewNotEqualsStr(), NewLikeIgnoreCase(), NewSet(), NewLikeList(), NewLikeInList())
	l.defaultOperands = make([]interface{}, 0)
	l.operandMap = make(map[string]interface{}, 0)
	l.operatorMap = make(map[string]interface{}, 0)
	l.operators = make([]Operator, 0)
	l.operands = make([]Operand, 0)
	l.operandList = make([]interface{}, 0)
	l.operatorList = make([]interface{}, 0)
	l.operatorIdentifiers = make(map[string]string, 0)

	for _, op := range l.defaultOperators {
		l.RegisterOperator(op.(Operator))
	}

	if l.registerDefaultOperands {
		for _, oper := range l.defaultOperands {
			l.registerOperand(oper.(Operand), false)
		}
	}

	l.setFieldSeparator("$")
	l.initialized = true

	return nil
}

func (l *LanguageImpl) setFieldSeparator(fieldSeparator string) {
	l.fieldSeparator = fieldSeparator
}

// RegisterOperand register operands with the language
func (l *LanguageImpl) RegisterOperand(op Operand) error {
	return l.registerOperand(op, true)
}

func (l *LanguageImpl) registerOperand(op Operand, errorIfAlreadyRegistered bool) error {
	name := op.GetName()
	_, found := l.operandMap[name]
	if found && errorIfAlreadyRegistered {
		return fmt.Errorf("Operand '%s' already registered", name)
	}
	l.operandMap[name] = op
	l.operandList = append(l.operandList, op)
	l.operands = append(l.operands, op)
	log.Debugf("Registered operand '%s' in lanuguage '%s'\n", name, l.name)

	return nil
}

// RegisterOperator register an operator with the language
func (l *LanguageImpl) RegisterOperator(op Operator) error {
	return l.registerOperator(op)
}

func (l *LanguageImpl) registerOperator(op Operator) error {
	oper := op.(Operator)
	id := strings.ToLower(oper.getName())
	l.operatorIdentifiers[id] = id
	l.operatorMap[id] = oper
	l.operatorList = append(l.operatorList, oper)
	l.operators = append(l.operators, oper)

	log.Debugf("Registered operator '%s' in lanuguage '%s'\n", op.getInputElements(), l.name)
	return nil
}

// CreateBooleanExpression creates an boolean expression
func (l *LanguageImpl) CreateBooleanExpression(exprStr string) (BooleanExpression, error) {
	log.Infof("Creating boolean expression: %s\n", exprStr)
	expr, err := l.CreateExpression(exprStr, booleanType)
	if err != nil {
		return nil, fmt.Errorf("Error occured while creating boolean expression: %s", err)
	}
	expr = l.peel(expr)
	log.Debugf("Finished creating boolean expression: %s\n", exprStr)
	return expr.(BooleanExpression), nil
}

// CreateExpression creates an expression
func (l *LanguageImpl) CreateExpression(exprStr string, typ TypeImpl) (Expression, error) {
	lexer := NewLexer(exprStr, l.fieldSeparator)
	expr, err := l.parse(lexer, 0, typ)
	if err != nil {
		return nil, fmt.Errorf("Error occuring during parsing of string '%s': %s", exprStr, err)
	}
	return expr, nil
}

func (l *LanguageImpl) peel(expr Expression) Expression {
	return expr
}

func (l *LanguageImpl) parse(lex *Lexer, depth int, typ TypeImpl) (Expression, error) {
	operatorFound := false
	done := false
	var list *MyList
	elements := make([]interface{}, 0)
	for {
		token := lex.nextToken()
		log.Debugf("token: %s\n", lex.getString())
		switch token {
		case lparen:
			log.Debugf("LPAREN")
			expr, err := l.parse(lex, depth+1, typ)
			if err != nil {
				return nil, fmt.Errorf("Error occured during parsing: %s", err)
			}
			if expr != nil {
				elements = append(elements, expr)
			}
			break
		case rparen:
			log.Debugf("RPAREN")
			if depth <= 0 {
				return nil, errors.New("Too many close parenthesis")
			}
			if !operatorFound && (list == nil) {
				ec := NewExpressionContext(emptyListType.TypeImpl)
				list = NewMyList(*ec)
			}
			done = true
		case comma:
			log.Debugf("COMMA")
		case eof:
			log.Debugf("EOF")
			if depth > 0 {
				return nil, errors.New("Missing close parenthesis")
			}
			done = true
		case integer:
			log.Debugf("INTEGER")
			i, _ := lex.getInt()
			ic := NewIntConstant(i)
			elements = append(elements, ic)
		case literal:
			log.Debugf("LITERAL")
			strLiteral := NewStringLiteral(lex.getString())
			elements = append(elements, strLiteral)
		case str:
			log.Debugf("STRING")
			id := lex.getString()
			strs := split(id, lex.getFieldSeparator())
			op := l.getOperand(strs[0])
			if op != nil {
				names := op.(Operand).GetFieldNames()
				fieldValues := []string{strconv.Itoa(max(len(strs)-1, len(names)))}
				if len(fieldValues) > 0 {
					if len(strs) > 1 {
						// Handle old way of specifying arguments, positionally with the field separator
						// Example: operandName$arg1Value$arg2Value
						// In this case, the example string above was returned by the lexer.
						for i := 1; i < len(strs); i++ {
							fieldValues[i-1] = strs[i]
						}
					} else {
						// Handle new way of specifying arguments, like a method call
						// Example 1: operandName(arg1Value,'arg 2 value')
						// Example 2: operandName(arg2Name='arg2Value',arg1Name='arg1Value')
						// In this case, only "operandName" was returned by the lexer, so we
						// have to call the lexer to parse the parameter lists
						token := lex.nextToken()
						if token != lparen {
							lex.pushBack()
						} else {
							for idx := 0; ; idx++ {
								token = lex.nextToken()
								if token == rparen {
									break
								}
								if (token != str) && (token != literal) {
									return nil, fmt.Errorf("Parsing - Invalid argument: %s", lex.getExprStr())
								}
								firstStr := lex.getString()
								token = lex.nextToken()
								if token == comma {
									fieldValues[idx] = firstStr
								} else if token == rparen {
									fieldValues[idx] = firstStr
									break
								} else if token == str {
									equals := lex.getString()
									if equals == "=" {
										return nil, fmt.Errorf("Parsing - expecting ',', '=', or ')' but found %s", equals)
									}
									token = lex.nextToken()
									if (token != str) && (token != literal) {
										return nil, fmt.Errorf("Parsing - invalid argument value in %s", lex.getExprStr())
									}
									fieldName := firstStr
									fieldValue := lex.getString()
									fieldIdx := 0
									for fieldIdx < len(names) {
										if strings.ToLower(fieldName) == names[fieldIdx].(Identifier).getName() {
											break
										}
										fieldIdx++
									}
									if fieldIdx >= len(names) {
										if len(names) == 0 {
											return nil, fmt.Errorf("Parsing - invalid expression: %s" + lex.getExprStr())
										}
										return nil, fmt.Errorf("Parsing - '%s' is an unknown argument name", fieldName)
									}
									fieldValues[fieldIdx] = fieldValue
									token = lex.nextToken()
									if token == rparen {
										break
									}
									if token != comma {
										return nil, fmt.Errorf("Parsing - '%s' is an unknown argument name; expression: %s", fieldName, lex.getExprStr())
									}
								} else {
									return nil, fmt.Errorf("Parsing - expecting ',', '=', or ')' in '%s'" + lex.getExprStr())
								}
							}
						}
					}
				}
				// Conver fieldvalues into an interface array
				b := make([]interface{}, len(fieldValues))
				for i := range fieldValues {
					b[i] = fieldValues[i]
				}
				info := newOperandInfo(op.(Operand), b)
				elements = append(elements, info)
				log.Debugf("operand= %s", info.toString())
			} else if l.isOperatorIdentifier(strings.ToLower(id)) {
				elements = append(elements, id)
				log.Debugf("operator= %s", id)
			} else {
				return nil, errors.New("Parsing - Unknown language ID")
			}
		default:
			return nil, errors.New("Parsing - unknown token")
		}
		// Try to reduce
		if l.reduce(&elements) {
			operatorFound = true
		}
		if token == comma {
			if len(elements) != 1 {
				return nil, errors.New("List element not resolved")
			}
			expr, err := l.getExpression(elements, nil)
			if err != nil {
				return nil, fmt.Errorf("Error occured while getting expression from elements: %s", err)
			}
			log.Debugf("Comma expr= %s", expr)
			if list == nil {
				ec := NewExpressionContext(expr.getType())
				list = NewMyList(*ec)
			}
			list.AddExpression(expr)
			elements = elements[:0] // Clear elements list

		} else if done {
			log.Debugf("Done at depth= %d", depth)
			break
		}
	}
	expr, err := l.getExpression(elements, &typ)
	if err != nil {
		return nil, fmt.Errorf("Error occured while getting expression from elements: %s", err)
	}
	log.Debugf("GetExpression - expression: %s\n", expr.toString())
	if list != nil {
		if expr != nil {
			log.Debugf("Added expr to list")
			list.AddExpression(expr)
		}
		expr = list.CreateListExpression(ExpressionContext{
			Typ: list.listType.TypeImpl,
		}).(Expression)
		log.Debugf("Created new expression ", expr, " from list ", list)
	}

	return expr, nil
}

func (l *LanguageImpl) getExpression(elements []interface{}, typ *TypeImpl) (Expression, error) {
	if len(elements) == 0 {
		return nil, nil
	}
	if len(elements) != 1 {
		return nil, errors.New("Unable to reduce expression")
	}
	expression, isExpression := elements[0].(Expression)
	if isExpression {
		return expression, nil
	}
	operandInformation, isOperand := elements[0].(Operand)
	if typ != nil && isOperand {
		fN := make([]interface{}, len(operandInformation.GetFieldNames()))
		fieldNames := operandInformation.GetFieldNames()
		for i, arg := range fieldNames {
			fN[i] = arg.(interface{})
		}
		ctx := NewOperandContext(typ, fN, l)
		return operandInformation.CreateExpression(*ctx), nil
	}
	return nil, errors.New("Not an expression")
}

func (l *LanguageImpl) reduce(elements *[]interface{}) bool {
	reduced := false
	printElements(*elements, "reduce")
	for l.reduce2(elements) {
		reduced = true
	}
	return reduced
}

func (l *LanguageImpl) reduce2(elements *[]interface{}) bool {
	for i := 0; i < len(l.operators); i++ {
		operator := l.operators[i]
		log.Debugf("operator= %s, signature= %+v, return= %s\n", operator.getName(), operator.getInputElements(), operator.getReturnType())
		for j := 0; j < len(*elements); j++ {
			expr := l.createExpression(elements, operator, j)
			if expr != nil {
				log.Debugf("Matched operator: %s", operator)

				// Remove elements
				for len(*elements) > j {
					if len(*elements) > 0 {
						*elements = (*elements)[:len(*elements)-1]
					}
				}
				*elements = append(*elements, expr)
				return true
			}
		}
	}
	return false
}

func (l *LanguageImpl) createExpression(elements *[]interface{}, oper Operator, start int) Expression {
	log.Debugf("createExpression - start: %d", start)
	expectedEles := oper.getInputElements()
	// var operand Operand
	if (len(*elements) - start) != len(expectedEles) {
		return nil
	}
	var input []interface{}
	// Iterate through the expected types for the operator
	// If there is no match, return null;
	// otherwise, build up the input list for this operator.
	for i := 0; i < len(expectedEles); i++ {
		expectedEle := expectedEles[i]
		actualEle := (*elements)[i+start]
		log.Debugf("Expected element: %s, actual element: %s", printElement(expectedEle), printElement(actualEle))
		typ, isType := expectedEle.(Type)
		var isIntType bool
		if isType {
			isIntType = reflect.DeepEqual(typ.getType(), intType)
		}
		_, isStringLiteral := actualEle.(StringLiteral)
		operandInformation := &OperandImpl{}
		isOperandInfo := false
		// Convert from StringLiteral to int if there is a previous operand of type 'int'
		// with selectable values of type 'int'.
		// For example, in the expression:
		//      (dayofweek >= 'Monday') and (dayOfWeek <= 'Friday')
		// 'dayofweek' is an operand that returns an int, and both 'Monday' and 'Friday' are
		// string literals that are converted to an int.
		if ((i > 0) && expectedEle == isIntType) && (isStringLiteral) {
			// Search for the first operand looking backwards in the elements list.
			for j := i - 1; j >= 0; j-- {
				ele := (*elements)[j]
				operandInformation, isOperandInfo = ele.(*OperandImpl)
				if isOperandInfo {
					break
				}
			}
			// If we found an operand
			if operandInformation != nil {
				strLiteral := actualEle.(StringLiteral)
				vals := operandInformation.GetSelectableValues()
				if len(vals) > 0 {
					for j := 0; j < len(vals); j++ {
						v := &SelectableValueImpl{}
						v = vals[j].(*SelectableValueImpl)
						if v.hasIntValue() {
							if v.getName() == strLiteral.getName() {
								actualEle = NewIntConstant(v.getIntValue())
								break
							}
						}
					}
				}
				return nil
			}
		}
		// Expected elements for an operator must be either:
		// 	(1) a "String" to denote the name of the operator.  Note that a single operator can consist of
		//      multiple strings.  For example,
		//             <int1> between <int2> and <int3>
		//      contains 2 string literals ('between' and 'and') but is a single operator.
		//  (2) a Type
		expectedStr, isString := expectedEle.(string)
		_, isStringArray := expectedEle.([]string)
		if isString || isStringArray {
			actualStr, isStringActualElem := actualEle.(string)
			if !isStringActualElem {
				return nil
			}
			if isString {
				if strings.ToLower(expectedStr) != strings.ToLower(actualStr) {
					return nil
				}
			}
		} else {
			_, isType := expectedEle.(TypeImpl)
			if !isType {
				_, isListType := expectedEle.(ListType)
				if !isListType {
					return nil
				}
			}
			expectedEleType := expectedEle.(Type)
			actualExpr, isExpr := actualEle.(Expression)
			if isExpr {
				actualType := actualExpr.getType()
				log.Debugf("Actual type: %s", actualType)
				log.Debugf("Expected type: %s", expectedEleType.getType())
				if !reflect.DeepEqual(*expectedEleType.getType(), actualType) {
					return nil
				}
				input = append(input, actualEle)
			} else {
				actualOperand, isOperand := actualEle.(*operandInfo)
				if isOperand {
					o := actualOperand.op
					expType := expectedEle.(Type)
					if !matchesType(*expType.getType(), o.GetReturnTypes()) {
						return nil
					}
					fieldNames := o.GetFieldNames()
					b := make([]interface{}, len(fieldNames))
					for i := range fieldNames {
						b[i] = fieldNames[i]
					}
					operandCtx := NewOperandContext(expectedEle.(Type), b, l)
					expr := o.CreateExpression(*operandCtx)
					input = append(input, expr)
				}
			}
		}
	}
	ctx := NewOperatorContext(oper, input, oper.getReturnType(), l)
	return oper.createExpression(*ctx)
}

func matchesType(typ TypeImpl, types []TypeImpl) bool {
	for i := 0; i < len(types); i++ {
		if typ.toString() == types[i].toString() {
			return true
		}
	}
	return false
}

func (l *LanguageImpl) isOperatorIdentifier(id string) bool {
	_, ok := l.operatorIdentifiers[id]
	return ok
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (l *LanguageImpl) getOperand(name string) interface{} {
	return l.operandMap[name]
}

func split(input, sep string) []string {
	fields := make([]string, 0)
	fieldStart := 0
	scanStart := 0
	for {
		idx := strings.Index(input, sep)
		if idx < 0 {
			fields = append(fields, input[fieldStart:])
			break
		}
		if idx > scanStart {
			if (string(input[idx-1]) == "\\") && ((idx == (scanStart + 1)) || (string(input[idx-2]) != "\\")) {
				input = input[:idx-1] + input[idx:]
				scanStart = idx + len(sep) - 1
				continue
			}
		}
		fields = append(fields, input[fieldStart:idx])
		scanStart = idx + len(sep)
		fieldStart = scanStart
	}
	return fields
}

// // ObtainEvaluationContext ...
// func (l *LanguageImpl) ObtainEvaluationContext(object interface{}) EvaluationContextImpl {
// 	var ctx EvaluationContextImpl
// 	size := len(l.ctxList)

// 	if size > 0 {
// 		if len(l.ctxList) > 0 {
// 			ctx = l.ctxList[size].(EvaluationContextImpl)
// 			l.ctxList = l.ctxList[:size-1]
// 		}
// 	}
// 	return ctx
// }

// // ReleaseEvaluationContext ...
// func (l *LanguageImpl) ReleaseEvaluationContext(evaluationContext EvaluationContext) {

// }

// Lexer ...
type Lexer struct {
	expr  string
	chars []rune
	index int
	// sb             *bytes.Buffer
	sb             []string
	str            string
	lastToken      int
	pushedBack     bool
	fieldSeparator string
}

const (
	literalDelimeter = `\'`
	none             = -1
	eof              = 0
	integer          = 1
	str              = 3
	literal          = 4
	lparen           = 5
	rparen           = 6
	comma            = 7
)

// NewLexer returns an instance of Lexer
func NewLexer(exprStr, fieldSeparator string) *Lexer {
	log.Debugf("Creating instance of lexer to lex string '%s'\n", exprStr)
	lexer := &Lexer{}
	lexer.init(exprStr, fieldSeparator)
	return lexer
}

func (lx *Lexer) init(expr string, fieldSeparator string) {
	lx.expr = expr
	lx.chars = []rune(expr)
	lx.sb = make([]string, 0)
	lx.fieldSeparator = fieldSeparator
}

func (lx *Lexer) nextToken() int {
	if lx.pushedBack {
		lx.pushedBack = false
		return lx.lastToken
	}
	for {
		log.Debugf("lx.index: %d, lx.chars: %d\n", lx.index, len(lx.chars))
		if lx.index >= len(lx.chars) {
			return eof
		}
		ch := lx.chars[lx.index]
		lx.index = lx.index + 1
		if isWhiteSpace(ch) {
			continue
		}
		// If we're looking at a (i.e. 'ch' is a) single quote, then skip past it and
		// return the literal between the quotes.  Note that index always points one character
		// beyond what ch has in it.
		if isLiteralDelimiter(ch) {
			lastCharWasEscape := false
			for lx.index < len(lx.chars) {
				ch = lx.chars[lx.index]
				lx.index = lx.index + 1
				if isLiteralDelimiter(ch) && !lastCharWasEscape {
					return lx.getToken(literal)
				}
				if lastCharWasEscape {
					lx.sb = append(lx.sb, "\\")
				}
				lastCharWasEscape = string(ch) == "\\"
				if !lastCharWasEscape {
					lx.sb = append(lx.sb, string(ch))
				}
			}
		}
		lx.sb = append(lx.sb, string(ch))
		if isSpecial(ch) {
			return lx.getToken(str)
		}

		if isDigit(ch) {
			for lx.index < len(lx.chars) {
				ch = lx.chars[lx.index]
				if !isDigit(ch) {
					break
				}
				lx.sb = append(lx.sb, string(ch))
				lx.index = lx.index + 1
			}
			return lx.getToken(integer)
		}
		if isStringStart(ch) {
			containsFieldSeparator := false
			escape := false
			for lx.index < len(lx.chars) {
				ch = lx.chars[lx.index]
				if escape {
					lx.sb = append(lx.sb, string(ch))
					escape = false
				} else {
					escape = string(ch) == "\\"
					if !escape {
						if !isStringPart(ch, containsFieldSeparator) {
							break
						}
						lx.sb = append(lx.sb, string(ch))
					} else {
						if (lx.index+1) < len(lx.chars) && (string(lx.chars[lx.index+1]) == "\\") {
							lx.index = lx.index + 1
						}
					}
				}
				if !containsFieldSeparator {
					containsFieldSeparator = stringInSlice(lx.sb, lx.getFieldSeparator())
				}
				lx.index = lx.index + 1
			}
			return lx.getToken(str)
		}
		for lx.index < len(lx.chars) {
			ch = lx.chars[lx.index]
			if isDigit(ch) || isStringStart(ch) || isLiteralDelimiter(ch) || isSpecial(ch) || isWhiteSpace(ch) {
				break
			}
			lx.sb = append(lx.sb, string(ch))
			lx.index = lx.index + 1
		}
		return lx.getToken(str)
	}
}

func stringInSlice(str []string, a string) bool {
	for _, b := range str {
		if b == a {
			return true
		}
	}
	return false
}

func (lx *Lexer) getFieldSeparator() string {
	return lx.fieldSeparator
}

func (lx Lexer) setFieldSeparator(fieldSeparator string) {
	lx.fieldSeparator = fieldSeparator
}

func isStringPart(ch rune, expandedChars bool) bool {
	if expandedChars {
		return unicode.IsLetter(ch) || string(ch) == "." || (string(ch) == "-") || (string(ch) == "*") || (string(ch) == "/") || (string(ch) == ":") || (string(ch) == "(") || (string(ch) == ")") || (string(ch) == ",") || (string(ch) == "[") || (string(ch) == "]") || (string(ch) == "=") || (string(ch) == "<") || (string(ch) == ">")
	}
	return unicode.IsLetter(ch) || string(ch) == "." || (string(ch) == "-")
}

func isLiteralDelimiter(ch rune) bool {
	if strings.ContainsRune(literalDelimeter, ch) {
		return true
	}
	return false
}

func isWhiteSpace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func isSpecial(ch rune) bool {
	return (string(ch) == "(" || string(ch) == ")")
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isStringStart(ch rune) bool {
	return unicode.IsLetter(ch)
}

func (lx *Lexer) getToken(token int) int {
	lx.str = strings.Join(lx.sb, "")
	lx.sb = lx.sb[:0]
	if token == str {
		if lx.str == "(" {
			token = lparen
		} else if lx.str == ")" {
			token = rparen
		} else if lx.str == "," {
			token = comma
		}
	}
	lx.lastToken = token
	return token
}

func (lx *Lexer) getInt() (int, error) {
	return strconv.Atoi(lx.getString())
}

func (lx *Lexer) getString() string {
	return lx.str
}

func (lx *Lexer) pushBack() {
	lx.pushedBack = true
}

func (lx *Lexer) getExprStr() string {
	return lx.expr
}

// MyList ...
type MyList struct {
	ListExpression

	exprList    []interface{}
	elementType *TypeImpl
	constant    bool
	listType    ListType
}

// NewMyList ...
func NewMyList(ctx ExpressionContext) *MyList {
	ml := &MyList{}

	ml.ListExpression = *NewListExpression(ctx)
	ml.listType = emptyListType

	return ml
}

// Size ...
func (ml *MyList) Size() int {
	return len(ml.exprList)
}

// GetExpression ...
func (ml *MyList) GetExpression(index int) Expression {
	return ml.exprList[index].(Expression)
}

// AddExpression ...
func (ml *MyList) AddExpression(expression interface{}) error {
	expr := expression.(Expression)
	typ := expr.getType()
	if ml.elementType != nil {
		if !reflect.DeepEqual(typ, *ml.elementType) {
			return errors.New("Elements mismatch")
		}
	} else {
		ml.elementType = &typ
		typeFound := ml.listType.findByElementType(*ml.elementType)
		if typeFound == nil {
			ml.listType = *NewListType(*ml.elementType)

		} else {
			ml.listType = *NewListType(*typeFound.elementType)
			ec := expr.getExpressionContext()
			ec.setType(ml.listType.TypeImpl)
		}
	}
	if !expr.isConstant() {
		ml.constant = false
	}
	ml.exprList = append(ml.exprList, expr)
	return nil
}

// CreateListExpression ...
func (ml *MyList) CreateListExpression(context ExpressionContext) interface{} {
	if reflect.DeepEqual(ml.listType, emptyListType) {
		return new(EmptyList)
	} else if reflect.DeepEqual(ml.listType.elementType, stringListType.elementType) {
		return newStringList(context, ml.exprList, ml.constant)
	} else if reflect.DeepEqual(ml.listType, intListType) {
		return newIntList(context, ml.exprList, ml.constant)
	} else if reflect.DeepEqual(ml.listType, booleanListType) {
		return newBooleanList(context, ml.exprList, ml.constant)
	}

	return nil
}

// EmptyList ...
type EmptyList struct {
	ListExpression
}

// ToString ...
func (el *EmptyList) toString() string {
	return "()"
}

type stringList struct {
	StringArrayExpressionImpl
	exprs []StringExpression
	vals  []string
}

func newStringList(context ExpressionContext, exprList []interface{}, constant bool) *stringList {
	sl := &stringList{}

	sl.StringArrayExpressionImpl.ExpressionImpl.Context = context
	sl.exprs = make([]StringExpression, len(exprList))
	for i, expr := range exprList {
		sl.exprs[i] = expr.(StringExpression)
	}
	if constant {
		sl.vals = sl.getVals(nil)
	} else {
		sl.vals = nil
	}

	return sl
}

func (sl *stringList) evaluate(ctx EvaluationContext) []string {
	if len(sl.vals) > 0 {
		return sl.vals
	}
	return sl.getVals(ctx)
}

func (sl *stringList) getVals(ctx EvaluationContext) []string {
	vals := make([]string, len(sl.exprs))

	for i := range sl.exprs {
		vals[i] = sl.exprs[i].evaluate(ctx)
	}

	return vals
}

func (sl *stringList) toString() string {
	var strArray []string
	for _, expr := range sl.exprs {
		strArray = append(strArray, expr.toString())
	}
	return strings.Join(strArray, ",")
}

type intList struct {
	*IntArrayExpressionImpl
	exprs []IntExpression
	vals  []int
}

func newIntList(context ExpressionContext, exprList []interface{}, constant bool) *intList {
	il := &intList{}

	il.IntArrayExpressionImpl = NewIntArrayExpression(context)
	il.exprs = make([]IntExpression, len(exprList))
	for i, expr := range exprList {
		il.exprs[i] = expr.(IntExpression)
	}
	if constant {
		il.vals = il.getVals(nil)
	} else {
		il.vals = nil
	}

	return il
}

func (il *intList) evaluate(ctx EvaluationContext) []int {
	if len(il.vals) > 0 {
		return il.vals
	}
	return il.getVals(ctx)
}

func (il *intList) getVals(ctx EvaluationContext) []int {
	vals := make([]int, len(il.exprs))

	for i := range il.exprs {
		vals[i] = il.exprs[i].evaluate(ctx)
	}

	return vals
}

func (il *intList) toString() string {
	var strArray []string
	for _, expr := range il.exprs {
		strArray = append(strArray, expr.toString())
	}
	return strings.Join(strArray, ",")
}

type booleanList struct {
	*BooleanArrayExpressionImpl
	exprs []BooleanExpression
	vals  []bool
}

func newBooleanList(context ExpressionContext, exprList []interface{}, constant bool) *booleanList {
	bl := &booleanList{}

	bl.BooleanArrayExpressionImpl = NewBooleanArrayExpression(context)
	bl.exprs = make([]BooleanExpression, len(exprList))
	for i, expr := range exprList {
		bl.exprs[i] = expr.(BooleanExpression)
	}
	if constant {
		bl.vals = bl.getVals(nil)
	} else {
		bl.vals = nil
	}

	return bl
}

func (bl *booleanList) evaluate(ctx EvaluationContext) []bool {
	if len(bl.vals) > 0 {
		return bl.vals
	}
	return bl.getVals(ctx)
}

func (bl *booleanList) getVals(ctx EvaluationContext) []bool {
	vals := make([]bool, len(bl.exprs))

	for i := range bl.exprs {
		vals[i] = bl.exprs[i].evaluate(ctx)
	}

	return vals
}

type operandInfo struct {
	op          Operand
	fieldvalues []interface{}
}

func newOperandInfo(op Operand, fieldValues []interface{}) *operandInfo {
	oi := &operandInfo{}

	oi.op = op
	oi.fieldvalues = fieldValues

	return oi
}

func (oi *operandInfo) toString() string {
	return oi.op.GetName()
}

type print interface {
	toString() string
}

func printElements(elements []interface{}, prefix string) {
	log.Debugf("<<<<< Printing elements: <<<<<<<<<<<")
	for _, elem := range elements {
		stringType, isString := elem.(string)
		stringArrayType, isStringArray := elem.([]string)
		if isString {
			log.Debugf("%s - Elements are: %+v\n", prefix, stringType)
		} else if isStringArray {
			log.Debugf("%s - Elements are: %+v\n", prefix, stringArrayType)
		} else {
			printElem := elem.(print)
			log.Debugf("%s - Elements are: %s\n", prefix, printElem.toString())
		}

	}
	log.Debugf(">>>>> Done Printing elements >>>>>>>")
}

func printElement(element interface{}) string {
	stringType, isString := element.(string)
	stringArrayType, isStringArray := element.([]string)
	if isString {
		return stringType
	} else if isStringArray {
		str := strings.Join(stringArrayType, ",")
		return str
	}

	printElem := element.(print)
	return printElem.toString()
}
