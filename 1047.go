// +build 1047

package goleet

// 删除字符串中的所有相邻重复项
// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string
func removeDuplicates(S string) string {
	r := make([]byte, 0, len(S)/2)

	for _, v := range []byte(S) {
		if len(r) == 0 || r[len(r)-1] != v {
			r = append(r, v)
		} else {
			r = r[:len(r)-1]
		}
	}

	return string(r)
}
