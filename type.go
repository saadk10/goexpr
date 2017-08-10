package expr

// Type defines the interface
type Type interface {
	getType() *TypeImpl
}

// TypeImpl implements Type
type TypeImpl struct {
	name string
}

var (
	booleanType = TypeImpl{"boolean"}
	intType     = TypeImpl{"int"}
	longType    = TypeImpl{"long"}
	stringType  = TypeImpl{"string"}
	emptyType   = TypeImpl{""}
)

// NewType returns an instance of TypeImpl
func NewType(name string) *TypeImpl {
	return &TypeImpl{
		name: name,
	}
}

func (t TypeImpl) toString() string {
	return t.name
}

func (t TypeImpl) getType() *TypeImpl {
	return &t
}
