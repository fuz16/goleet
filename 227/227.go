package leet

// 基本计算器 II
// https://leetcode-cn.com/problems/basic-calculator-ii/
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		stackNum []int
		stackOp  []byte
	)

	value := 0
	for _, v := range []byte(s + "$" ) {
		if v == ' ' {
			continue
		}

		if c := int(v - '0'); c >= 0 && c <= 9 {
			value = value*10 + c
			continue
		}
		stackNum = append(stackNum, value)
		value = 0
		if len(stackOp) == 0 {
			stackOp = append(stackOp, v)
			continue
		}
	LOOP:
		for len(stackOp) > 0 && len(stackNum) > 1 {
			switch op := stackOp[len(stackOp)-1]; op {
			case '+', '-':
				if v == '*' || v == '/' {
					break LOOP
				}

				if op == '+' {
					stackNum[len(stackNum)-2] += stackNum[len(stackNum)-1]
				}

				if op == '-' {
					stackNum[len(stackNum)-2] -= stackNum[len(stackNum)-1]
				}

			case '*':
				stackNum[len(stackNum)-2] *= stackNum[len(stackNum)-1]
			case '/':
				stackNum[len(stackNum)-2] /= stackNum[len(stackNum)-1]
			}
			stackOp = stackOp[:len(stackOp)-1]
			stackNum = stackNum[:len(stackNum)-1]
		}

		stackOp = append(stackOp, v)
	}

	if len(stackNum) > 0 {
		return stackNum[0]
	}

	return 0
}
