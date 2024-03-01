package resource

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func ConvertK8sResourceFile(fname string) {
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
		obj, gvk, err := decoder.Decode([]byte(resourceYAML), nil, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decode resource YAML")
			continue
		}

		log.Print(gvk)

		switch {
		case (gvk.Group == "apps" && gvk.Version == "v1" && gvk.Kind == "Deployment"):
			UpdateDeployment(obj.(*appsv1.Deployment)) //nolint:forcetypeassert // known type
			log.Print(obj.(*appsv1.Deployment))        //nolint:forcetypeassert // known type
		// TODO: cases for other resource types
		default:
			log.Error().
				Dict("gvk", zerolog.Dict().
					Str("group", gvk.Group).
					Str("version", gvk.Version).
					Str("kind", gvk.Kind)).
				Msg("Resource is not a supported type")
		}
		// TODO: output to files
	}
}
