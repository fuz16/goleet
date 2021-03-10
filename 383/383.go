package leet

// 赎金信
// https://leetcode-cn.com/problems/ransom-note/
func canConstruct(ransomNote string, magazine string) bool {
	statistics := make([]int, 26)

	for _, v := range []byte(magazine) {
		statistics[v-'a']++
	}

	for _, v := range []byte(ransomNote) {
		if statistics[v-'a'] == 0 {
			return false
		}
		statistics[v-'a']--
	}

	return true
}
