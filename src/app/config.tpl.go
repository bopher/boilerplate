package app

import "github.com/bopher/config"

// SetupConfig driver
func SetupConfig() {
	// {{if eq .config "env"}}
	if c, ok := config.NewEnvConfig("./config/config.env"); ok {
		_container.Register("--APP-CONFIG", c)
	} else {
		panic("failed to build config driver")
	}
	// {{end}}
	// {{if eq .config "json"}}
	if c, ok := config.NewJSONConfig("./config/config.json"); ok {
		_container.Register("--APP-CONFIG", c)
	} else {
		panic("failed to build config driver")
	}
	// {{end}}
	// {{if eq .config "memory"}}
	if c, ok := config.NewMemoryConfig(nil); ok {
		_container.Register("--APP-CONFIG", c)
	} else {
		panic("failed to build config driver")
	}
	// {{end}}
}

// Config get config manager
// leave name empty to resolve default
func Config(names ...string) config.Config {
	name := "--APP-CONFIG"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(config.Config); ok {
			return res
		}
	}
	return nil
}
