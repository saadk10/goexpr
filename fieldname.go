package expr

// FieldName defines the interface
type FieldName interface {
	getSelectableValues() []SelectableValue
	getDefaultValues() string
	validateValue(value string)
}

// FieldNameImpl implements FieldName
type FieldNameImpl struct {
	selectableValues []SelectableValue
}
