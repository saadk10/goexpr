package expr

// Identifier defines the interface
type Identifier interface {
	getName() string
}

// IdentifierImpl implements Identifier
type IdentifierImpl struct {
	name, key string
}

// NewIdentifier returns an instance of Identifer
func NewIdentifier(name string) *IdentifierImpl {
	return &IdentifierImpl{
		name: name,
		key:  "",
	}
}

// NewIdentifierNameKey returns an instance of Identifer
func NewIdentifierNameKey(name, key string) *IdentifierImpl {
	return &IdentifierImpl{
		name: name,
		key:  key,
	}
}

func (i *IdentifierImpl) getName() string {
	return i.name
}

func (i *IdentifierImpl) toString() string {
	return i.name
}
