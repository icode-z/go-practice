package main

import (
	"fmt"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
// 输出: true
// 示例 2:
//
// 输入: "()[]{}"
// 输出: true
// 示例 3:
//
// 输入: "(]"
// 输出: false
// 示例 4:
//
// 输入: "([)]"
// 输出: false
// 示例 5:
//
// 输入: "{[]}"
// 输出: true

var symbolMap = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	symbol := make([]rune, 0)
	length := len(s)
	if length%2 != 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			symbol = append(symbol, rune(s[i]))
			continue
		}

		// 查询出栈的值是否对应
		if _d, ok := symbolMap[rune(s[i])]; ok {
			if _d == symbol[len(symbol)-1] {
				if len(symbol) == 1 {
					return true
				}
				symbol = symbol[:len(symbol)-1]
				continue
			}
		}
	}
	return false
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
}
