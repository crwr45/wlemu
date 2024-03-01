package updaters

import (
	"fmt"

	"github.com/crwr45/wlemu/pkg/rule"
	"github.com/rs/zerolog/log"
	corev1 "k8s.io/api/core/v1"
)

// ContainerUpdateFunc is the interface an Updater MUST provide as the entrypoint.
// When passed it will be passed the base container definition and the raw (un-decoded)
// configuration for the rule that matched it to the container definition.
// It must return a string containing YAML defining any otherresources the new container
// requires. If an error occurs it MUST return an empty string and an error.
type ContainerUpdateFunc func(*corev1.Container, rule.RawNode) (string, error)

type ResourceBuilder struct {
	UpdateContainer ContainerUpdateFunc
	Name            string
}

var registry = make(map[string]*ResourceBuilder)

func Register(identifier string, containerBuildFunc ContainerUpdateFunc) error {
	if _, exists := registry[identifier]; exists {
		err := fmt.Errorf("cannot register %s, already registered", identifier)
		log.Error().Err(err).Msg("")
		return err
	}
	log.Info().Msgf("registering Updater '%s'", identifier)
	registry[identifier] = &ResourceBuilder{
		Name:            identifier,
		UpdateContainer: containerBuildFunc,
	}

	return nil
}

func GetUpdater(name string) (*ResourceBuilder, error) {
	if rb, exists := registry[name]; !exists {
		return nil, fmt.Errorf("no updater found for name '%s'", name)
	} else {
		return rb, nil
	}
}
