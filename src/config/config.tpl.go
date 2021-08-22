package config

import "github.com/bopher/config"

// Configure register/override app config
func Configure(config config.Config) {
	config.Set("name", "// {{.name}}")
	config.Set("locale", "// {{.locale}}")
	config.Set("key", "// {{.appKey}}")
	// {{if eq .database "mongo"}}
	config.Set("mongo.conStr", "mongodb://localhost:27017")
	// {{end}}
	// {{if eq .database "mysql"}}
	config.Set("mysql.host", "")
	config.Set("mysql.username", "root")
	config.Set("mysql.password", "root")
	// {{end}}
}
