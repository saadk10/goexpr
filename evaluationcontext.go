package expr

// EvaluationContext defines the interface
type EvaluationContext interface {
	GetObject() interface{}
}

// EvaluationContextImpl implements EvaluationContext
type EvaluationContextImpl struct {
	object interface{}
	index  map[string]interface{}
}

// GetObject returns the object
func (e *EvaluationContextImpl) GetObject() interface{} {
	return e.object
}

func (e *EvaluationContextImpl) setObject(object interface{}) {
	e.object = object
}

func (e *EvaluationContextImpl) setIndex(index map[string]interface{}) {
	e.index = index
}

func (e *EvaluationContextImpl) getProperty(key string) interface{} {
	return e.index[key]
}

func (e *EvaluationContextImpl) setProperty(key string, property interface{}) {
	e.index[key] = property
}

func (e *EvaluationContextImpl) clear() {
	e.object = nil
	e.index = make(map[string]interface{})
}
