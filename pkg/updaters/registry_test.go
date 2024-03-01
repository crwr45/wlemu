package updaters_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	rule "github.com/crwr45/wlemu/pkg/rule"
	"github.com/crwr45/wlemu/pkg/updaters"
	corev1 "k8s.io/api/core/v1"
)

const identifier = "test_identifier"

func BuildContainerTestFunc(container *corev1.Container, rawConf rule.RawNode) (string, error) {
	return "", nil
}

var _ = Describe("Registry", Ordered, func() {
	When("registering a new entry", func() {
		It("should register a new entry without error", func() {
			err := updaters.Register(identifier, BuildContainerTestFunc)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should error if the entry has already been registered", func() {
			err := updaters.Register(identifier, BuildContainerTestFunc)
			Expect(err).To(HaveOccurred())
		})
	})
})
