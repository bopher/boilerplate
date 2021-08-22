package app

import (
	"github.com/bopher/translator"
)

// SetupTranslator driver
func SetupTranslator() {
	// {{if eq .translator "json"}}
	if t, err := translator.NewJSONTranslator("// {{.locale}}", "./config/strings"); err == nil {
		_container.Register("--APP-TRANSLATOR", t)
	} else {
		panic("failed to build translator driver")
	}
	// {{else}}
	if t := translator.NewMemoryTranslator("// {{.locale}}"); t != nil {
		_container.Register("--APP-TRANSLATOR", t)
	} else {
		panic("failed to build translator driver")
	}
	// {{end}}
}

// Translator get translator driver
// leave name empty to resolve default
func Translator(names ...string) translator.Translator {
	name := "--APP-TRANSLATOR"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(translator.Translator); ok {
			return res
		}
	}
	return nil
}
