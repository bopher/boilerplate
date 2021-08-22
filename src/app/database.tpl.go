package app

import (
	migrations "github.com/bopher/cliutils/migration"
	"github.com/bopher/database"
	"github.com/jmoiron/sqlx"
)

// SetupDatabase driver
func SetupDatabase() {
	conf := Config()
	if db, err := database.NewMySQLConnector(
		conf.String("mysql.host", ""),
		conf.String("mysql.username", "root"),
		conf.String("mysql.password", ""),
		conf.String("database.name", "// {{.name}}"),
	); err == nil {
		_container.Register("--APP-DB", db)
	} else {
		panic("failed to init database " + err.Error())
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
