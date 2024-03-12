package rule

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// TODO consider top-level sections for different types of resource in the future
type RuleEntry struct {
	Config        RawNode `yaml:"config"`
	ContainerName string  `yaml:"container_name"`
	Replacement   string  `yaml:"replacement"`
}

// LoadRulesFromFile is the entrypoint for the rule package; it will load a rule
// file, populate the top-level struct but keep raw config for individual rules
func LoadRulesFromFile(fname string) {
	log.Info().Msgf("Loading rules from file: %s", fname)

	contents, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read Rule file")
	}

	var rules = []RuleEntry{}
	err = yaml.Unmarshal(contents, &rules)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load Rule file")
	}

	for _, rule := range rules {
		err := registerRule(rule)
		if err != nil {
			log.Error().Err(err).Msg("failed to register Rule")
			continue
		}
	}
}

type Rule struct {
	Config      RawNode
	UpdaterName string
}

var rules = make(map[string]*Rule)

func registerRule(newRule RuleEntry) error {
	if _, exists := rules[newRule.ContainerName]; exists {
		return fmt.Errorf("cannot register for container name '%s', already registered", newRule.ContainerName)
	}
	rules[newRule.ContainerName] = &Rule{
		Config:      newRule.Config,
		UpdaterName: newRule.Replacement,
	}
	log.Debug().Str("rule_name", newRule.ContainerName).Str("updater", newRule.Replacement).Msg("registered Rule")
	return nil
}

// This is trivial to have a function for but rule matching will almost certainly need to
// be much more expressive and flexible; e.g. using regular expressions or wildcards
func GetRuleForContainerName(name string) (*Rule, error) {
	if r, exists := rules[name]; !exists {
		return nil, fmt.Errorf("no rule found for name '%s'", name)
	} else {
		return r, nil
	}
}
