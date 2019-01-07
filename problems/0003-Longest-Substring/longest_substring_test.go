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

var _ = Describe("", func() {
	Describe("lengthOfLongestSubstring1", func() {
		Context("with basic cases", func() {
			It("returns correct results", func() {
				Expect(
					lengthOfLongestSubstring1("abcabcbb")).
					To(Equal(3))
				Expect(
					lengthOfLongestSubstring1("bbbbb")).
					To(Equal(1))
				Expect(lengthOfLongestSubstring1("pwwkew")).
					To(Equal(3))
			})
		})

		Context("with more difficult cases", func() {
			It("returns reccect results", func() {
				Expect(
					lengthOfLongestSubstring1("abcdefghijklmnopqrstuvwxyzABCD")).
					To(Equal(30))
			})
		})
	})

	Describe("lengthOfLongestSubstring2", func() {
		Context("with basic cases", func() {
			It("returns correct results", func() {
				Expect(
					lengthOfLongestSubstring2("abcabcbb")).
					To(Equal(3))
				Expect(
					lengthOfLongestSubstring2("bbbbb")).
					To(Equal(1))
				Expect(lengthOfLongestSubstring2("pwwkew")).
					To(Equal(3))
				Expect(
					lengthOfLongestSubstring2("abcdefghijklmnopqrstuvwxyzABCD")).
					To(Equal(30))
			})
		})
	})

	Describe("lengthOfLongestSubstring3", func() {
		Context("with basic cases", func() {
			It("returns correct results", func() {
				Expect(
					lengthOfLongestSubstring3("abcabcbb")).
					To(Equal(3))
				Expect(
					lengthOfLongestSubstring3("bbbbb")).
					To(Equal(1))
				Expect(lengthOfLongestSubstring3("pwwkew")).
					To(Equal(3))
				Expect(
					lengthOfLongestSubstring3("abcdefghijklmnopqrstuvwxyzABCD")).
					To(Equal(30))
			})
		})
	})

	Describe("lengthOfLongestSubstring4", func() {
		Context("with basic cases", func() {
			It("returns correct results", func() {
				Expect(
					lengthOfLongestSubstring4("abcabcbb")).
					To(Equal(3))
				Expect(
					lengthOfLongestSubstring4("bbbbb")).
					To(Equal(1))
				Expect(lengthOfLongestSubstring4("pwwkew")).
					To(Equal(3))
				Expect(
					lengthOfLongestSubstring4("abcdefghijklmnopqrstuvwxyzABCD")).
					To(Equal(30))
			})
		})
	})
})
