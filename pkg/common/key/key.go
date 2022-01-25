package key

import (
	"github.com/alisyahbana/efishery-test/pkg/common/config"
	"github.com/alisyahbana/efishery-test/pkg/common/env"
)

type Config struct {
	SignatureJwt string `json:"signatureJwt"`
}

var cfg *Config

func init() {
	config.LoadConfiguration(&cfg, "key", env.GetEnv())
}

func GetConfig() Config {
	return *cfg
}
