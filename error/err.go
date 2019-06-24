package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(value float64)(float64, error) {
	if(value < 0){
		// errors.New来构造一个基本的错误消息
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func main() {


	if result, err:= Sqrt(-1); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	if result, err := Sqrt(16); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
}


