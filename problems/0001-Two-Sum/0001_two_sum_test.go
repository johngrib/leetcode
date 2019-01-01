package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var _ = Describe("twoSum", func() {
	Context("with a basic case", func() {
		given := []int{2, 7, 11, 15}
		target := 9

		It("returns indices of 2 and 7", func() {
			result := twoSum(given, target)

			Expect(result).To(Equal([]int{0, 1}))
		})
	})
	Context("with a custom case", func() {
		given := []int{3, 3}
		target := 6

		It("returns indices of 3 and 3", func() {
			result := twoSum(given, target)

			Expect(result).To(Equal([]int{0, 1}))
		})
	})
})
