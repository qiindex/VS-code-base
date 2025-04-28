package go_demo

import "testing"

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"空字符串", "", true},
		{"简单匹配", "()", true},
		{"嵌套匹配", "({[]})", true},
		{"不匹配", "(]", false},
		{"顺序错误", "([)]", false},
		{"只有左括号", "({[", false},
		{"只有右括号", ")}]", false},
		{"混合匹配", "()[]{}", true},
		{"复杂嵌套", "({[()]})", true},
		{"错误嵌套", "({[()]}", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidParentheses(tt.s); got != tt.want {
				t.Errorf("IsValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidParenthesesOptimized(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"空字符串", "", true},
		{"简单匹配", "()", true},
		{"嵌套匹配", "({[]})", true},
		{"不匹配", "(]", false},
		{"顺序错误", "([)]", false},
		{"只有左括号", "({[", false},
		{"只有右括号", ")}]", false},
		{"混合匹配", "()[]{}", true},
		{"复杂嵌套", "({[()]})", true},
		{"错误嵌套", "({[()]}", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidParenthesesOptimized(tt.s); got != tt.want {
				t.Errorf("IsValidParenthesesOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkIsValidParentheses(b *testing.B) {
	s := "({[()]})" * 1000 // 创建一个长字符串
	for i := 0; i < b.N; i++ {
		IsValidParentheses(s)
	}
}

func BenchmarkIsValidParenthesesOptimized(b *testing.B) {
	s := "({[()]})" * 1000 // 创建一个长字符串
	for i := 0; i < b.N; i++ {
		IsValidParenthesesOptimized(s)
	}
}
