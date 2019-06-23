 package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	// switch的使用：不需要加break，每个case后面默认有break
	//g := ""
	switch {
	//case score < 0 || score > 100:
	//	// panic: 中断程序的执行
	//	panic(fmt.Sprintf(
	//		"Wrong score: %d", score))
	case score < 60 && score >= 0:
		//g = "F"
		return "F"
	case score < 80 && score > 0:
		//g = "C"
		return "C"
	case score < 90 && score > 0:
		//g = "B"
		return "B"
	case score <= 100 && score > 0:
		//g = "A"
		return "A"
	default:
		return  "输入错误"
	}
	//return g
}

func main() {

	const (
		filename = "abc.txt"
		)
	// ioutil.Readfile读取文件
	// nil:空值
	// if的使用：if可以像for一样写，加；分号
	// contents和err是在if里面定义的，出了if语句，contents和err都不能访问
	if contents, err := ioutil.ReadFile(filename); err != nil {
		//如果错误不为空，则打印错误
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//a, b := ioutil.ReadFile(filename)
	//if b !=nil {
	//	fmt.Print(b)
	//}else {
	//	fmt.Print(a)
	//}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		// Uncomment to see it panics.
		grade(-3),
	)
}
