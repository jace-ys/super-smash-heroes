package integration

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuperheroService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SuperheroService Suite")
}
