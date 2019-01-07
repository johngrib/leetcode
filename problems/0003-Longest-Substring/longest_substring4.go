package main

// sliding window optimized using array
func lengthOfLongestSubstring4(s string) int {
	str := []rune(s)
	length := len(str)
	set := [128]rune{}

	result, left, right := 0, 0, 0

	for ; right < length; right++ {
		char := str[right]
		left = max(int(set[char]), left)
		result = max(result, right-left+1)
		set[char] = rune(right + 1)
	}
	return result
}
