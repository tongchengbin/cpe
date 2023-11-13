package cpe

import (
	"fmt"
	"regexp"
)

type CPE struct {
	Part      string
	Vendor    string
	Product   string
	Version   string
	Update    string
	Edition   string
	Language  string
	SwEdition string
	TargetSw  string
	TargetHw  string
	Other     string
}

func Parse(cpeString string) (*CPE, error) {
	// 检查 CPE 字符串的版本信息
	if isCPE23(cpeString) {
		return ParseCPE23(cpeString)
	} else if isCPE22(cpeString) {
		return ParseCPE22(cpeString)
	} else {
		return nil, fmt.Errorf("无法识别的CPE版本: %s", cpeString)
	}
}
func ParseCPE23(cpeString string) (*CPE, error) {
	// 添加 CPE 2.3 版本的解析逻辑
	// 示例正则表达式
	cpeRegex := regexp.MustCompile(`^cpe:2\.3:([aho]?):([*a-zA-Z0-9._%+-]+):([*a-zA-Z0-9._%+-]+):([*a-zA-Z0-9._%+-]+)(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?(?::([*a-zA-Z0-9._%+-]+|[*]))?$`)
	// 使用正则表达式匹配 CPE 字符串
	matches := cpeRegex.FindStringSubmatch(cpeString)
	if matches == nil {
		return nil, fmt.Errorf("CPE字符串无效: %s", cpeString)
	}
	// 创建并填充 CPE 结构体
	cpe := &CPE{
		Part:      matches[1],
		Vendor:    "2.3",
		Product:   matches[3],
		Version:   matches[4],
		Update:    matches[5],
		Edition:   matches[6],
		Language:  matches[7],
		SwEdition: matches[8],
		TargetSw:  matches[9],
		TargetHw:  matches[10],
		Other:     matches[11],
	}

	return cpe, nil
}

func ParseCPE22(cpeString string) (*CPE, error) {
	// 添加 CPE 2.2 版本的解析逻辑
	// 示例正则表达式
	cpeRegex := regexp.MustCompile(`^cpe:/([aho])?:([a-zA-Z0-9._%+-]+):([a-zA-Z0-9._%+-]+):([a-zA-Z0-9._%+-]+)(?::([a-zA-Z0-9._%+-]+))?(?::([a-zA-Z0-9._%+-]+))?(?::([a-zA-Z0-9._%+-]+))?(?::([a-zA-Z0-9._%+-]+))?(?::([a-zA-Z0-9._%+-]+))?$`)

	// 使用正则表达式匹配 CPE 字符串
	matches := cpeRegex.FindStringSubmatch(cpeString)
	if matches == nil {
		return nil, fmt.Errorf("CPE字符串无效: %s", cpeString)
	}

	// 创建并填充 CPE 结构体
	cpe := &CPE{
		Part:      matches[1],
		Vendor:    matches[2],
		Product:   matches[3],
		Version:   matches[4],
		Update:    matches[5],
		Edition:   matches[6],
		Language:  matches[7],
		SwEdition: matches[8],
		TargetSw:  matches[9],
		TargetHw:  matches[10],
		Other:     matches[11],
	}

	return cpe, nil
}

func main() {
	// 测试用的 CPE 字符串
	cpeString := "cpe:2.3:a:appserv_open_project:appserv:*:*:*:*:*:*:*:*"

	// 解析 CPE 字符串
	cpe, err := Parse(cpeString)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	// 打印解析结果
	fmt.Println("Part:", cpe.Part)
	fmt.Println("Vendor:", cpe.Vendor)
	fmt.Println("Product:", cpe.Product)
	fmt.Println("Version:", cpe.Version)
	fmt.Println("Update:", cpe.Update)
	fmt.Println("Edition:", cpe.Edition)
	fmt.Println("Language:", cpe.Language)
	fmt.Println("SwEdition:", cpe.SwEdition)
	fmt.Println("TargetSw:", cpe.TargetSw)
	fmt.Println("TargetHw:", cpe.TargetHw)
	fmt.Println("Other:", cpe.Other)
}
