package app

import (
	ccrypto "github.com/bopher/cliutils/crypto"
	"github.com/bopher/crypto"
)

// SetupCrypto driver
func SetupCrypto() {
	conf := Config()
	if c := crypto.NewCryptography(conf.String("key", "")); c != nil {
		_container.Register("--APP-CRYPTO", c)
	} else {
		panic("failed to build crypto driver")
	}
	_cli.AddCommand(ccrypto.HashCommand(func(driver string) crypto.Crypto {
		return Crypto(driver)
	}, "--APP-CRYPTO"))
}

// Crypto get crypto driver
// leave name empty to resolve default
func Crypto(names ...string) crypto.Crypto {
	name := "--APP-CRYPTO"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(crypto.Crypto); ok {
			return res
		}
	}
	return nil
}
