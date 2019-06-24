package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 字符串转byte数组，len(s)：获得字节长度，而不是字符数
	s := "Yes我爱BUPT!" // UTF-8
	fmt.Println([]byte(s), len(s))


	for _, b := range []byte(s) {
		// %X:16进制显示
		// 英文一字节，一个汉字占三个字节
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// utf8.RuneCountInString(s): 输入字符数，英文/中文都是一个字符
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		// utf8.DecodeRune(bytes): 将字符一个个取出
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 字符串转字符数组
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c %d) ", i, ch, len([]rune(s) ))
	}
	fmt.Println()
}
