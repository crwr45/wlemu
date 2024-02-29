package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type configFile struct {
	Replacements []Replacement `yaml:"replacements"`
}

type Replacement struct {
	Config      RawNode `yaml:"config"`
	Name        string  `yaml:"name"`
	Replacement string  `yaml:"replacement"`
}

var Config = configFile{}

// LoadConfigFromFile is the entrypoint for the config package; it will load a configuration
// file, populate the top-level config, and any custom config sections that have
// been registered prior to the config being loaded.
func LoadConfigFromFile(fname string) {
	log.Info().Msgf("Loading config from file: %s", fname)

	contents, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read Config file")
	}

	err = yaml.Unmarshal(contents, &Config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load Config file")
	}
}
