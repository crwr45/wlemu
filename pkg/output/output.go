package output

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/runtime"
	k8sJson "k8s.io/apimachinery/pkg/runtime/serializer/json"
)

func SerializeToFile(obj runtime.Object, extras []string, fname string) {
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0660) //nolint:gomnd // file perms
	if err != nil {
		log.Error().Err(err).Str("fname", fname).Msg("failed to open file")
	}
	defer file.Close()

	log.Info().Str("fname", fname).Msg("writing updated object to file")

	if len(extras) > 0 {
		log.Debug().Str("fname", fname).Any("gvk", obj.GetObjectKind().GroupVersionKind()).Msg("found extras to prepend")
		joined := strings.Join(extras, "\n---\n") + "\n---\n"
		_, err = file.WriteString(joined)
		if err != nil {
			log.Error().Err(err).Str("fname", fname).Msg("failed to write extras to file")
		}
	}

	k8sSerializer := k8sJson.NewSerializerWithOptions(
		k8sJson.DefaultMetaFactory, nil, nil,
		k8sJson.SerializerOptions{
			Yaml:   true,
			Pretty: true,
			Strict: true,
		},
	)
	err = k8sSerializer.Encode(obj, file)
	if err != nil {
		log.Error().Err(err).Str("fname", fname).Any("gvk", obj.GetObjectKind().GroupVersionKind()).Msg("failed to serialize")
	}
}
