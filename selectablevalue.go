package expr

// SelectableValue defines the interface
type SelectableValue interface {
	getFieldNames() []FieldName
}

// SelectableValueImpl implements SelectableValue
type SelectableValueImpl struct {
	IdentifierImpl
	hasIntVal  bool
	intValue   int
	fieldNames []FieldName
}

func (svi *SelectableValueImpl) getFieldNames() []FieldName {
	return svi.fieldNames
}

func (svi *SelectableValueImpl) hasIntValue() bool {
	return svi.hasIntVal
}

func (svi *SelectableValueImpl) getIntValue() int {
	return svi.intValue
}
