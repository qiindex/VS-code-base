package go_demo

// IsValidParentheses 判断括号是否有效
// 使用栈的数据结构来实现
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func IsValidParentheses(s string) bool {
	// 使用map存储括号的对应关系
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 使用切片模拟栈
	stack := make([]rune, 0)

	// 遍历字符串
	for _, char := range s {
		// 如果是右括号
		if matching, isRight := pairs[char]; isRight {
			// 如果栈为空或栈顶元素不匹配，返回false
			if len(stack) == 0 || stack[len(stack)-1] != matching {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈
			stack = append(stack, char)
		}
	}

	// 如果栈为空，说明所有括号都匹配
	return len(stack) == 0
}

// IsValidParenthesesOptimized 优化版的有效括号判断
// 在遇到右括号时直接判断，不需要存储左括号
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func IsValidParenthesesOptimized(s string) bool {
	// 使用计数器来跟踪括号的嵌套层级
	var count1, count2, count3 int

	for _, char := range s {
		switch char {
		case '(':
			count1++
		case ')':
			count1--
			if count1 < 0 {
				return false
			}
		case '[':
			count2++
		case ']':
			count2--
			if count2 < 0 {
				return false
			}
		case '{':
			count3++
		case '}':
			count3--
			if count3 < 0 {
				return false
			}
		}
	}

	// 所有计数器都应该为0
	return count1 == 0 && count2 == 0 && count3 == 0
}
