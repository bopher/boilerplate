package app

import (
	"path"

	"github.com/bopher/cache"
	"github.com/bopher/config"
	"github.com/bopher/crypto"
	"github.com/bopher/logger"

	// {{if eq .database "mysql"}}
	"github.com/jmoiron/sqlx"
	// {{end}}
)

func confOrPanic() config.Config {
	if Config() == nil {
		panic("Failed to find default config")
	}
	return Config()
}

func cacheOrPanic() cache.Cache {
	if Cache() == nil {
		panic("Failed to find default cache")
	}
	return Cache()
}

// ConfigPath get configs path
func ConfigPath(sub ...string) string {
	return path.Join(append([]string{".", "config"}, sub...)...)
}

// LogPath get logs path
func LogPath(sub ...string) string {
	return path.Join(append([]string{".", ".logs"}, sub...)...)
}

// StoragePath get storage path
func StoragePath(sub ...string) string {
	return path.Join(append([]string{".", ".storage"}, sub...)...)
}

// DateFormatter get default app date formatter
func DateFormatter() logger.TimeFormatter {
	// {{if eq .locale "fa"}}
	return logger.JalaliFormatter
	// {{else}}
	return logger.GregorianFormatter
	// {{end}}
}

// IsUnderMaintenance check if under maintenance mode
func IsUnderMaintenance() (bool, error) {
	return cacheOrPanic().Exists("maintenance")
}

// CryptoResolver resolve crypto driver by name
func CryptoResolver(driver string) crypto.Crypto {
	return Crypto(driver)
}

// CacheResolver resolve cache driver by name
func CacheResolver(driver string) cache.Cache {
	return Cache(driver)
}

// {{if eq .database "mysql"}}
// DatabaseResolver resolve database driver by name
func DatabaseResolver(driver string) *sqlx.DB {
	return Database(driver)
}

// {{end}}
