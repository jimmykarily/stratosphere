package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStratosphere(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stratosphere Suite")
}
