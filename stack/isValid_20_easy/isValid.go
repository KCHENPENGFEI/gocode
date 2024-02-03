package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/13 下午2:52
 * @Desc:
 */

func main() {

}

func isValid(s string) bool {
	match := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if match[top] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
