package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	// for循环的使用
	//把整数转换成二进制表达式
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		//strconv: 字符转换包，itoa数字转字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	//os.open打开文件
	file, err := os.Open(filename)
	if err != nil {
		//panic把程序停下来报错
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	//逐行读取文件
	scanner := bufio.NewScanner(reader)

	//相当于while，go里面没有while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	//死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	fmt.Println("abc.txt contents:")
	printFile("basic/branch/abc.txt")

	fmt.Println("printing a string:")
	s := `abc"d"
    kkkk
	123

	p`
	printFileContents(strings.NewReader(s))

	// Uncomment to see it runs forever
	// forever()
}
