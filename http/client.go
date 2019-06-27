package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) " +
					"AppleWebKit/602.1.50 (KHTML, like Gecko) " +
						"CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	//resp, err := http.DefaultClient.Do(request)
	/**

	 */
	client := http.Client{
		//如果CheckRedirect不为nil，客户端会在执行重定向之前调用本函数字段。
		//参数req和via是将要执行的请求和已经执行的请求（切片，越新的请求越靠后）。
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	/**
		函数调用结束之前response.body.Close()一定要关闭
	 */
	defer resp.Body.Close()

	/**
		httputil.DumpResponse(resp, true):转储respone的body到内存中
	 */
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("======================")
	fmt.Printf("%s\n", s)
}
