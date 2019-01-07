package main

// sliding window
func lengthOfLongestSubstring2(s string) int {

	str := []rune(s)
	length := len(str)
	set := make(map[rune]bool)

	result, left, right := 0, 0, 0

	for left < length && right < length {
		char := str[right]

		if set[char] {
			set[str[left]] = false
			left++

		} else {
			set[char] = true
			right++
			result = max(result, right-left)
		}
	}

	return result
}
