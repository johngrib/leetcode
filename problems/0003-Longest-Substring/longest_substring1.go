package main

// brute force - O(n^3)
func lengthOfLongestSubstring1(s string) int {

	length := len(s)
	result := 0

	for start := 0; start < length; start++ {
		for end := start + 1; end <= length; end++ {
			if isAllUnique(s[start:end]) {
				result = max(result, end-start)
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isAllUnique(sub string) bool {
	set := make(map[rune]bool)
	chars := []rune(sub)

	for _, char := range chars {
		if set[char] {
			return false
		}
		set[char] = true
	}
	return true
}

func main() {
	lengthOfLongestSubstring1("pwwkew")
}
