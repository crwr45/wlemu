package resource

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func LoadK8sResourceFile(fname string) {
	// Load the file into a buffer
	data, err := os.ReadFile(fname)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to open %s", fname)
	}

	// Create a runtime.Decoder from the Codecs field within
	// k8s.io/client-go that's pre-loaded with the schemas for all
	// the standard Kubernetes resource types.
	decoder := scheme.Codecs.UniversalDeserializer()

	for _, resourceYAML := range strings.Split(string(data), "---") {
		// skip empty documents, `Decode` will fail on them
		if resourceYAML == "" {
			continue
		}

		// - obj is the API object (e.g., Deployment)
		// - gvk is a generic object that allows
		//   detecting the API type we are dealing with, for
		//   accurate type casting later.
		obj, gvk, err := decoder.Decode(
			[]byte(resourceYAML),
			nil,
			nil)
		if err != nil {
			log.Print(err)
			continue
		}

		if gvk.Group == "apps" &&
			gvk.Version == "v1" &&
			gvk.Kind == "Deployment" {
			deployment := obj.(*appsv1.Deployment) //nolint:forcetypeassert // known type

			log.Print(deployment.Spec.Template.Spec.Containers[0].Name)
			log.Print(deployment.Spec.Template.Spec.Containers[0].Image)

			// builder, _ := getRegistered("stressng") // TODO find from config via name
			// builder.containerBuilderFunc(ds.Spec.Template.Spec.Containers[0], config.Config.Replacements[0].Config)
		}
	}
}
