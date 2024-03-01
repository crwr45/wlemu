package resource

import (
	"strings"

	"github.com/crwr45/wlemu/pkg/rule"
	"github.com/crwr45/wlemu/pkg/updaters"
	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
)

// Extract container definitions from the Deployment object, find a matching rule, and update the container definition
func UpdateDeployment(deployment *appsv1.Deployment) string {
	log.Debug().Str("deployment", deployment.Name).Msg("updating containers in Deployment")

	var extraResources = []string{}

	for idx := 0; idx < len(deployment.Spec.Template.Spec.Containers); idx++ {
		container := &deployment.Spec.Template.Spec.Containers[idx]
		log.Print(container.Name)
		matchedRule, err := rule.GetRuleForContainerName(container.Name)
		if err != nil {
			log.Info().Str("container_name", container.Name).Msg("no rule matches container")
			continue
		}
		if updater, err := updaters.GetUpdater(matchedRule.UpdaterName); err == nil {
			log.Print(container.Command)
			log.Info().Str("updater", updater.Name).Str("container", container.Name).Msg("updating container")
			extraRes, err := updater.UpdateContainer(container, matchedRule.Config)
			log.Print(container.Command)
			if err == nil {
				extraResources = append(extraResources, extraRes)
			}
		} else {
			log.Warn().Str("updater_name", matchedRule.UpdaterName).Msg("no updater found")
		}
	}
	return strings.Join(extraResources, "\n---\n")
}
