package app

import (
	migrations "github.com/bopher/cliutils/migration"
	"github.com/bopher/database"
	"github.com/jmoiron/sqlx"
)

// SetupDatabase driver
func SetupDatabase() {
	conf := Config()
	host := conf.Cast("mysql.host").StringSafe("")
	username := conf.Cast("mysql.username").StringSafe("root")
	password := conf.Cast("mysql.password").StringSafe("")
	db := conf.Cast("database.name").StringSafe("// {{.name}}")

	if db, err := database.NewMySQLConnector(host, username, password, db); err == nil {
		_container.Register("--APP-DB", db)
	} else {
		panic("failed to init mysql database: " + err.Error())
	}

	_cli.AddCommand(migrations.MigrationCommand(func(driver string) *sqlx.DB {
		return Database(driver)
	}, "--APP-DB", "./database/migrations", "./database/seeds"))
}

// Database get database driver
// leave name empty to resolve default
func Database(names ...string) *sqlx.DB {
	name := "--APP-DB"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(*sqlx.DB); ok {
			return res
		}
	}
	return nil
}
