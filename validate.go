package cpe

import "regexp"

// isCPE23 验证 CPE 2.3 版本的字符串是否有效
func isCPE23(cpeString string) bool {
	// 添加 CPE 2.3 版本的验证逻辑
	// 示例正则表达式
	//cpe:2.3:a:appserv_open_project:appserv:*:*:*:*:*:*:*:*
	cpeRegex := regexp.MustCompile(`^cpe:2\.3:([aho]?):[\p{Han}*a-zA-Z0-9._%+-]+:[\p{Han}*a-zA-Z0-9._%+-]+:[*a-zA-Z0-9._%+-]+(?::[*a-zA-Z0-9._%+-]+)*$`)
	return cpeRegex.MatchString(cpeString)

}

// isCPE22 验证 CPE 2.2 版本的字符串是否有效
func isCPE22(cpeString string) bool {
	// 添加 CPE 2.2 版本的验证逻辑
	// 示例正则表达式
	cpeRegex := regexp.MustCompile(`^cpe:[a-zA-Z0-9._%+-]+:[a-zA-Z0-9._%+-]+:[a-zA-Z0-9._%+-]+(:[a-zA-Z0-9._%+-]+)?(:[a-zA-Z0-9._%+-]+)?(:[a-zA-Z0-9._%+-]+)?(:[a-zA-Z0-9._%+-]+)?(:[a-zA-Z0-9._%+-]+)?$`)
	return cpeRegex.MatchString(cpeString)
}
