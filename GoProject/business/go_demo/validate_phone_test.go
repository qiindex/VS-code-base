package go_demo

import (
	"testing"
)

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"13812345678", true},   // 合法
		{"19912345678", true},   // 合法
		{"12345678901", false},  // 非法，第二位不是3-9
		{"1381234567", false},   // 非法，长度不足
		{"23812345678", false},  // 非法，首位不是1
		{"138123456789", false}, // 非法，长度超出
	}
	for _, tt := range tests {
		if got := ValidatePhone(tt.input); got != tt.expected {
			t.Errorf("ValidatePhone(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name+tag@domain.co", true},
		{"user@sub.domain.com", true},
		{"user@domain", false},    // 没有顶级域名
		{"user@.com", false},      // 域名不合法
		{"@domain.com", false},    // 缺少用户名
		{"userdomain.com", false}, // 缺少@符号
	}
	for _, tt := range tests {
		if got := ValidateEmail(tt.input); got != tt.expected {
			t.Errorf("ValidateEmail(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
