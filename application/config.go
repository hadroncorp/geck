package application

import (
	"github.com/caarlos0/env/v11"

	"github.com/hadroncorp/geck/versioning"
)

// Config configuration structure for most applications.
type Config struct {
	// ApplicationName name of the application.
	ApplicationName string `env:"NAME"`
	// Version the version tag of the application. Preferably with semantic version format.
	Version string `env:"VERSION"`
	// Environment running environment type (e.g. development, local, staging, production).
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
	// Semver the semantic version structure. Will be empty if Version has an invalid semver.
	Semver versioning.SemanticVersion
}

// NewConfig allocates a new Config instance.
// Parses Config.Version as versioning.SemanticVersion if Config.Version has a valid semver format.
func NewConfig() (Config, error) {
	opts := env.Options{
		Prefix:          "APP_",
		RequiredIfNoDef: true,
	}
	cfg, err := env.ParseAsWithOptions[Config](opts)
	if err != nil {
		return Config{}, err
	}
	cfg.Semver, err = versioning.NewSemanticVersion(cfg.Version)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
