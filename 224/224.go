package leet

// 基本计算器
// https://leetcode-cn.com/problems/basic-calculator/
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = "(" + s + ")"
	stack := make([]int, 0)
	stack2 := make([]byte, 0)
	value := -1
LOOP:
	for _, v := range []byte(s) {
		if v == ' ' {
			continue
		}
		if isNum(v) {
			if value == -1 {
				value = int(v - '0')
			} else {
				value = value*10 + int(v-'0')
			}
			continue
		}

		if value >= 0 {
			stack = append(stack, value)
			value = -1
		}

		if v == '(' {
			stack2 = append(stack2, v)
			continue
		}

	LOOP2:
		for (len(stack2) != 0 && len(stack) > 1) || (len(stack) == 1 && len(stack2) > 0 && stack2[len(stack2)-1] != '+') {
			op := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			data := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch op {
			case '+':
				stack[len(stack)-1] += data
			case '-':
				stack2 = append(stack2, '+')
				stack = append(stack, -data)
			case '(':
				stack = append(stack, data)
				if v == ')' {
					continue LOOP
				}
				stack2 = append(stack2, op)
				break LOOP2
			}
		}
		stack2 = append(stack2, v)
	}

	if len(stack) > 0 {
		return stack[0]
	}
	return 0
}

func isNum(c byte) bool {
	return c != '(' && c != ')' && c != '+' && c != '-'
}
