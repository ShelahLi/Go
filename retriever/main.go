package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
			"name":   "lth",
			"course": "golang",
		})
}

/**
	接口间的组合：Retriever + Poster
 */
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever

	mockRetriever := mock.Retriever{
		Contents: "this is a fake baidu.com"}
	r = &mockRetriever
	inspect(r)

	/**
	有real.Retriever的方法传入的值是指针类型，所以要加&，取地址
	*/
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	/**
	判断r是否为*mock.Retriever类型
	*/
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))

}

/**
打印对象的内部信息
*/
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	/**
	%T: 输出r的类型；%v：输出r的值
	*/
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	/**
	r.(type):获取r的类型，相当于：T%
	*/
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
