package app

import (
	"github.com/bopher/validator"
	"github.com/bopher/validator/validations"
)

// SetupValidator driver
func SetupValidator() {
	if v := validator.NewValidator(Translator(), "// {{.locale}}"); v != nil {
		// {{if eq .locale "fa"}}
		validations.RegisterExtraValidations(v)
		// {{end}}
		_container.Register("--APP-VALIDATOR", v)
	} else {
		panic("failed to build validator driver")
	}
}

// Validator get validator driver
// leave name empty to resolve default
func Validator(names ...string) validator.Validator {
	name := "--APP-VALIDATOR"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(validator.Validator); ok {
			return res
		}
	}
	return nil
}
