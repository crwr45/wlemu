package stressng

import (
	"strconv"

	"github.com/crwr45/wlemu/pkg/rule"
	"github.com/crwr45/wlemu/pkg/updaters"
	"github.com/rs/zerolog/log"
	corev1 "k8s.io/api/core/v1"
)

const (
	id        = "stress-ng"
	imageName = "ghcr.io/abraham2512/fedora-stress-ng:master"
	baseCmd   = "stress-ng"
)

var (
	newProbe corev1.ExecAction = corev1.ExecAction{Command: []string{"ls"}}
)

type StressNGConfig struct {
	Stressor  string `yaml:"stressor"`
	Instances int    `yaml:"instances"`
}

func init() {
	_ = updaters.Register(id, UpdateContainer)
}

func UpdateContainer(container *corev1.Container, rawConf rule.RawNode) (string, error) {
	conf := StressNGConfig{}
	if err := rawConf.Decode(&conf); err != nil {
		log.Error().Err(err).Str("container", container.Name).Str("id", id).Msgf("Failed to decode config")
		return "", err //nolint:wrapcheck // WiP
	}
	updateProbe(container.LivenessProbe)
	updateProbe(container.ReadinessProbe)
	updateProbe(container.StartupProbe)
	updateImage(container, &conf)
	return "", nil
}

func updateImage(container *corev1.Container, conf *StressNGConfig) {
	container.Image = imageName
	container.Command = []string{baseCmd, "--" + conf.Stressor, strconv.Itoa(conf.Instances), "-v"}
}

func updateProbe(probe *corev1.Probe) {
	if probe != nil {
		probe.Exec = &newProbe
		probe.GRPC = nil
		probe.HTTPGet = nil
		probe.TCPSocket = nil
	}
}
