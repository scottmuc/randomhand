package randomhand_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRandomhand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Randomhand Suite")
}
