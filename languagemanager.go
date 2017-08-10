package expr

// LanguageManager defines the interface
type LanguageManager interface {
	CreateLanguage(name string) (Language, error)
}

// LanguageManagerImpl defines the struct
type LanguageManagerImpl struct{}

// NewLanguageManger returns an instance of LanguageManagerImpl
func NewLanguageManger() *LanguageManagerImpl {
	lmgr := &LanguageManagerImpl{}
	return lmgr
}

// CreateLanguage creates a new language
func (lm *LanguageManagerImpl) CreateLanguage(name string, debug bool) (Language, error) {
	return lm.createLanguage(name, true, debug)
}

func (lm *LanguageManagerImpl) createLanguage(name string, registerDefaultOperands bool, debug bool) (Language, error) {
	lang := &LanguageImpl{
		name: name,
		registerDefaultOperands: registerDefaultOperands,
		debug: debug,
	}
	lang.init()
	return lang, nil
}
