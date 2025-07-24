package go_demo

import (
	"regexp"
)

// ValidatePhone 验证中国大陆手机号是否合法
func ValidatePhone(phone string) bool {
	// 中国大陆手机号正则：1开头，第二位3-9，后面9位0-9，共11位
	var re = regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}

// ValidateEmail 验证email地址是否合法
func ValidateEmail(email string) bool {
	// 常见email正则表达式
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
