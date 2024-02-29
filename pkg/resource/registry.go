package resource

import (
	"fmt"

	"github.com/crwr45/wlemu/pkg/config"
	"github.com/rs/zerolog/log"
	corev1 "k8s.io/api/core/v1"
)

type ContainerBuilderFunc func(*corev1.Container, config.RawNode) (*corev1.Container, []interface{}, error)

type ResourceBuilder struct {
	containerBuilderFunc ContainerBuilderFunc
	name                 string
}

var registry = make(map[string]ResourceBuilder)

func Register(identifier string, containerBuildFunc ContainerBuilderFunc) error {
	log.Info().Msgf("Registering '%s'", identifier)
	if _, exists := registry[identifier]; exists {
		err := fmt.Errorf("cannot register %s, already registered", identifier)
		log.Error().Err(err).Msg("")
		return err
	}
	registry[identifier] = ResourceBuilder{
		name:                 identifier,
		containerBuilderFunc: containerBuildFunc,
	}

	log.Debug().Msgf("Registered identifier '%s' for [%T]%v", identifier, registry[identifier], registry[identifier])
	log.Debug().Msgf("Current state of registry %v", registry)
	return nil
}
