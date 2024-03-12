package resource

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/crwr45/wlemu/pkg/output"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func ConvertK8sResourceFile(fname, outDir string) {
	data, err := os.ReadFile(fname)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to open %s", fname)
	}

	decoder := scheme.Codecs.UniversalDeserializer()

	for idx, resourceYAML := range strings.Split(string(data), "---") {
		if resourceYAML == "" {
			continue
		}

		obj, gvk, err := decoder.Decode([]byte(resourceYAML), nil, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decode resource YAML")
			continue
		}
		outFile := buildOutputPath(fname, outDir, idx)

		switch {
		case (gvk.Group == "apps" && gvk.Version == "v1" && gvk.Kind == "Deployment"):
			extras := UpdateDeployment(obj.(*appsv1.Deployment)) //nolint:forcetypeassert // known type
			output.SerializeToFile(obj, extras, outFile)
		// TODO: cases for other resource types
		default:
			log.Error().
				Dict("gvk", zerolog.Dict().
					Str("group", gvk.Group).
					Str("version", gvk.Version).
					Str("kind", gvk.Kind)).
				Msg("Resource is not a supported type")
		}
	}
}

func buildOutputPath(fname, outDir string, idx int) string {
	ext := filepath.Ext(fname)
	nameNoExt := strings.TrimSuffix(filepath.Base(fname), ext)

	name := fmt.Sprintf("%s-%d%s", nameNoExt, idx, ext)
	return filepath.Join(outDir, name)
}
