package bitcal_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBitCal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BitCal Suite")
}
