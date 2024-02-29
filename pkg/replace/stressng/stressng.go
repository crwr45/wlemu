package stressng

import (
	"github.com/crwr45/wlemu/pkg/config"
	"github.com/crwr45/wlemu/pkg/resource"
	"github.com/rs/zerolog/log"
	corev1 "k8s.io/api/core/v1"
)

const identifier = "stressng"

type StressNGConfig struct {
	Stressor  string `yaml:"stressor"`
	Instances int    `yaml:"instances"`
}

func init() {
	_ = resource.Register(identifier, BuildContainer)
}

func BuildContainer(container *corev1.Container, rawConf config.RawNode) (*corev1.Container, []interface{}, error) {
	cfg := StressNGConfig{}
	if err := rawConf.Decode(&cfg); err != nil {
		log.Error().Err(err).Msgf("Failed to decode config for %s", identifier)
		return container, nil, err //nolint:wrapcheck // WiP
	}
	log.Print(cfg)
	log.Print(cfg.Instances)
	log.Print(cfg.Stressor)
	return container, nil, nil
}
