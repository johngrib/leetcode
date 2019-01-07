package main

// sliding window optimized
func lengthOfLongestSubstring3(s string) int {
	str := []rune(s)
	length := len(str)
	set := make(map[rune]int)

	result, left, right := 0, 0, 0

	for ; right < length; right++ {
		char := str[right]

		if lastIndex, ok := set[char]; ok {
			left = max(lastIndex, left)
		}
		result = max(result, right-left+1)
		set[char] = right + 1
	}
	return result
}
