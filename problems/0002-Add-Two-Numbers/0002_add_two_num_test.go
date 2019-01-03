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

var _ = Describe("addTwoNumbers", func() {
	Context("with a basic case", func() {
		l1 := &ListNode{Val: 2}
		l2 := &ListNode{Val: 4}
		l3 := &ListNode{Val: 3}
		l1.Next = l2
		l2.Next = l3

		r1 := &ListNode{Val: 5}
		r2 := &ListNode{Val: 6}
		r3 := &ListNode{Val: 4}
		r1.Next = r2
		r2.Next = r3

		It("returns the correct result", func() {
			result := addTwoNumbers(l1, r1)
			Expect(result.Val).To(Equal(7))
			Expect(result.Next.Val).To(Equal(0))
			Expect(result.Next.Next.Val).To(Equal(8))
		})
	})

	Context("with a custom case", func() {
		l1 := &ListNode{Val: 5}
		r1 := &ListNode{Val: 5}

		It("returns the correct result", func() {
			result := addTwoNumbers(l1, r1)
			Expect(result.Val).To(Equal(0))
			Expect(result.Next.Val).To(Equal(1))
		})
	})
})
