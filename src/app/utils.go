package app

import (
	"path"

	"github.com/bopher/config"
)

func confOrPanic() config.Config {
	if Config() == nil {
		panic("Failed to find default config")
	}
	return Config()
}

func ConfigPath(sub ...string) string {
	return path.Join(append([]string{".", "config"}, sub...)...)
}

func LogPath(sub ...string) string {
	return path.Join(append([]string{".", ".logs"}, sub...)...)
}

func StoragePath(sub ...string) string {
	return path.Join(append([]string{".", ".storage"}, sub...)...)
}
