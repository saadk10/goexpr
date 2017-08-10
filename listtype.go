package expr

import "fmt"

// ListType ...
type ListType struct {
	TypeImpl
	elementType *TypeImpl
}

var (
	booleanListType = *NewListType(booleanType)
	intListType     = *NewListType(intType)
	stringListType  = *NewListType(stringType)
	emptyListType   = *NewListType(emptyType)
	all             = []ListType{booleanListType, intListType, stringListType, emptyListType}
)

// NewListType returns an instance of ListType
func NewListType(elementType TypeImpl) *ListType {
	eT := TypeImpl{
		name: fmt.Sprintf("List of %s", elementType.name),
	}
	listType := &ListType{
		elementType: &elementType,
	}
	listType.TypeImpl = eT
	return listType
}

// GetElementType ...
func (lt *ListType) getElementType() *TypeImpl {
	return lt.elementType
}

// SetElementType ...
func (lt *ListType) setElementType(elementType TypeImpl) {
	lt.elementType = &elementType
}

// FindByElementType ...
func (lt *ListType) findByElementType(elementType TypeImpl) *ListType {
	for i := range all {
		if elementType == *all[i].elementType {
			return &all[i]
		}
	}
	return nil
}

func (lt ListType) getType() *TypeImpl {
	return &lt.TypeImpl
}
