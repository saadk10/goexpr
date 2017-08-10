package expr

// OperandInfo ...
type OperandInfo interface {
	GetFieldNames() []FieldName
	GetSelectableValues() []SelectableValue
	GetName() string
}

// OperandInfoImpl ...
type OperandInfoImpl struct {
	IdentifierImpl
	FieldNames       []FieldName
	SelectableValues []SelectableValue
}

// NewOperand return instance of OperandInfoImpl for specified operand name
func NewOperand(name string) *OperandInfoImpl {
	return NewOperandWithKeyAndFieldNames(name, "", nil, nil)
}

// NewOperandWithFieldNames returns instance of OperandInfoImpl for specified operand name and field names
func NewOperandWithFieldNames(name string, fieldNames []FieldName, selectableValues []SelectableValue) *OperandInfoImpl {
	return NewOperandWithKeyAndFieldNames(name, "", fieldNames, selectableValues)
}

// NewOperandWithKeyAndFieldNames returns instance of OperandInfoImpl for specified operand name, key, and field names
func NewOperandWithKeyAndFieldNames(name string, key string, fieldNames []FieldName, selectableValues []SelectableValue) *OperandInfoImpl {
	o := &OperandInfoImpl{}
	o.IdentifierImpl = *NewIdentifierNameKey(name, key)
	o.FieldNames = fieldNames
	o.SelectableValues = selectableValues
	return o
}

// GetFieldNames returns field names
func (o OperandInfoImpl) GetFieldNames() []FieldName {
	return o.FieldNames
}

// GetName returns name
func (o OperandInfoImpl) GetName() string {
	return o.IdentifierImpl.getName()
}

// GetSelectableValues returns selectable values
func (o OperandInfoImpl) GetSelectableValues() []SelectableValue {
	return o.SelectableValues
}

func (o *OperandInfoImpl) toString() string {
	return o.GetName()
}
