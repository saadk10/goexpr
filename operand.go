package expr

// Operand defines the interface
type Operand interface {
	GetName() string
	GetFieldNames() []FieldName
	GetSelectableValues() []SelectableValue
	GetReturnTypes() []TypeImpl
	CreateExpression(ctx OperandContext) Expression
	toString() string
}

// OperandImpl defines thes struct
type OperandImpl struct {
	Info        *OperandInfoImpl
	ReturnTypes []TypeImpl
}

// NewOperandWithType creates an instance OperandImpl with the specified operand info and type
func NewOperandWithType(info OperandInfo, returnType TypeImpl) *OperandImpl {
	return &OperandImpl{
		Info:        info.(*OperandInfoImpl),
		ReturnTypes: []TypeImpl{returnType},
	}
}

// NewOperandWithTypeList creates an instance OperandImpl with the specified operand info and type list
func NewOperandWithTypeList(info OperandInfo, returnType []TypeImpl) *OperandImpl {
	return &OperandImpl{
		Info:        info.(*OperandInfoImpl),
		ReturnTypes: returnType,
	}
}

// NewOperandWithNameType creates an instance OperandImpl with the specified operand name and type
func NewOperandWithNameType(name string, returnType TypeImpl) *OperandImpl {
	operandInfo := NewOperand(name)

	return &OperandImpl{
		Info:        operandInfo,
		ReturnTypes: []TypeImpl{returnType},
	}
}

// NewOperandWithNameTypeList creates an instance OperandImpl with the specified operand name and type list
func NewOperandWithNameTypeList(name string, returnType []TypeImpl) *OperandImpl {
	operandInfo := NewOperand(name)

	return &OperandImpl{
		Info:        operandInfo,
		ReturnTypes: returnType,
	}
}

// GetFieldNames returns field name
func (o *OperandImpl) GetFieldNames() []FieldName {
	return o.Info.GetFieldNames()
}

// GetName returns name of operand
func (o *OperandImpl) GetName() string {
	return o.Info.GetName()
}

// GetSelectableValues returns selectable value
func (o *OperandImpl) GetSelectableValues() []SelectableValue {
	return o.Info.GetSelectableValues()
}

// GetReturnTypes returns return types
func (o *OperandImpl) GetReturnTypes() []TypeImpl {
	return o.ReturnTypes
}

func (o *OperandImpl) validateInput() bool {
	return true
}

func (o *OperandImpl) toString() string {
	return o.Info.GetName()
}
