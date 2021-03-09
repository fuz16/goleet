package leet

// 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	statistics := make([]int, 26)

	for i := 0; i < len(s); i++ {
		statistics[s[i]-'a']++
		statistics[t[i]-'a']--
	}

	for _, v := range statistics {
		if v != 0 {
			return false
		}
	}

	return true
}
