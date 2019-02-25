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

var _ = Describe("Median of Two Sorted Arrays", func() {
	Describe("findMedianSortedArrays1", func() {
		Context("default cases", func() {
			It("passes", func() {
				Expect(
					findMedianSortedArrays1([]int{1, 3}, []int{2})).
					To(Equal(2.0))

				Expect(
					findMedianSortedArrays1([]int{1, 2}, []int{3, 4})).
					To(Equal(2.5))
			})
		})

		Context("small cases", func() {
			It("passes", func() {
				Expect(
					findMedianSortedArrays1([]int{1}, []int{1})).
					To(Equal(1.0))

				Expect(
					findMedianSortedArrays1([]int{0}, []int{1})).
					To(Equal(0.5))
			})
		})
	})
})
