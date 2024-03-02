package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/9/8 下午3:14
 * @Desc:
 */

func main() {
	s := "bb3[a2[c]b]"
	fmt.Println(decodeString1(s))
}

type stack struct {
	arr []string
}

func (s *stack) pop() string {
	data := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]
	return data
}

func (s *stack) push(str string) {
	s.arr = append(s.arr, str)
}

func (s *stack) len() int {
	return len(s.arr)
}

func newStack() stack {
	arr := make([]string, 0)
	s := stack{
		arr: arr,
	}
	return s
}

func decodeString(s string) string {
	stackNum := newStack()
	//stackBrackets := newStack()
	stackStr := newStack()
	//var decoded strings.Builder

	numStr := make([]byte, 0)
	charStr := make([]byte, 0)
	for i := 0; i < len(s); i++ {

		char := s[i]
		if char == '[' {
			stackNum.push(string(numStr))
			numStr = make([]byte, 0)
			stackStr.push(string(charStr))
			charStr = make([]byte, 0)
		} else if char == ']' {
			numStr := stackNum.pop()
			c := stackStr.pop()
			num, _ := strconv.Atoi(numStr)
			for i := 0; i < num; i++ {
				c += c
			}
			charStr = []byte(c)
		} else if char >= 'a' && char <= 'z' {
			charStr = append(charStr, char)
		} else if char >= '0' && char <= '9' {
			numStr = append(numStr, char)
		} else {
		}
	}
	return string(charStr)
}

func decodeString1(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	result := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			strStack = append(strStack, result)
			result = ""
			numStack = append(numStack, num)
			num = 0
		} else if char == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			result = str + strings.Repeat(result, count)
		} else {
			result += string(char)
		}
	}
	return result
}

// decodeString2
//
//	@Description: 二刷
//	@param s
//	@return string
func decodeString2(s string) string {
	var numStack []int
	var strStack []string

	curNum, curStr := 0, ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			curNum = curNum*10 + int(c-'0')
		} else if c == '[' {
			// 压入栈中，并且清空当前处理变量
			numStack = append(numStack, curNum)
			curNum = 0
			strStack = append(strStack, curStr)
			curStr = ""
		} else if c == ']' {
			// 处理数据出栈
			// 先弹出数字
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			// 在弹出字符串
			curStr = strStack[len(strStack)-1] + strings.Repeat(curStr, count)
			strStack = strStack[:len(strStack)-1]
		} else {
			curStr += string(c)
		}
	}
	return curStr
}
