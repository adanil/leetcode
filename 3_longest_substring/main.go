package main

func lengthOfLongestSubstring(s string) int {
	ans := 0

	l, r := 0, 0

	m := make(map[uint8]int)

	for r < len(s) {
		c := s[r]
		if cnt := m[c]; cnt == 0 {
			r++
			m[c]++
			ans = max(ans, r-l)
		} else {
			cl := s[l]
			m[cl]--
			l++
		}
	}

	return ans
}

func main() {

}
