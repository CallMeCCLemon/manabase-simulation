package scryfall_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScryfall(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scryfall Suite")
}
