package main

import (
	"bytes"
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/1/13 下午2:16
 * @Desc: 给你一个字符串 s 和一个整数 repeatLimit ，用 s 中的字符构造一个新字符串 repeatLimitedString ，使任何字母 连续 出现的次数都不超过 repeatLimit 次。你不必使用 s 中的全部字符。

返回 字典序最大的 repeatLimitedString 。

如果在字符串 a 和 b 不同的第一个位置，字符串 a 中的字母在字母表中出现时间比字符串 b 对应的字母晚，则认为字符串 a 比字符串 b 字典序更大 。如果字符串中前 min(a.length, b.length) 个字符都相同，那么较长的字符串字典序更大。
*/

func main() {
	s := "fffff"
	l := 2
	r := repeatLimitedString(s, l)
	fmt.Println(r)
}

func repeatLimitedString(s string, repeatLimit int) string {
	byteS := []byte(s)
	sort.Slice(byteS, func(i, j int) bool {
		return byteS[i] > byteS[j]
	})
	var stack []byte
	var reBuild bytes.Buffer
	var lastByte byte
	var cnt int

	for i := 0; i < len(byteS) || len(stack) > 0; i++ {
		// 优先检查栈中数据
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top != lastByte {
				reBuild.WriteByte(top)
				lastByte = top
				stack = stack[0 : len(stack)-1]
				cnt = 1
			} else {
				if cnt < repeatLimit {
					reBuild.WriteByte(top)
					cnt++
					stack = stack[0 : len(stack)-1]
				} else {
					// 不能出栈
					if i >= len(byteS) {
						stack = stack[0 : len(stack)-1]
					} else {
						break
					}
				}
			}
		}
		if i < len(byteS) {
			if byteS[i] != lastByte {
				reBuild.WriteByte(byteS[i])
				lastByte = byteS[i]
				cnt = 1
			} else {
				if cnt < repeatLimit {
					reBuild.WriteByte(byteS[i])
					lastByte = byteS[i]
					cnt++
				} else {
					// 不能继续加入后续，先存入栈中
					stack = append(stack, byteS[i])
				}
			}
		}

	}
	return reBuild.String()
}
