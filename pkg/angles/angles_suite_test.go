package angles_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAngles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Angles Suite")
}
