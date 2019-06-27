package main

import (
	"fmt"
	"math"
)

// 定义一个接口：图形接口
type Shape interface {
	area() float64
}

// 圆实现这个接口
type Circle struct {
	x, y, radius float64
}

// 长方形
type Rectangle struct {
	height, width float64
}

// 各自实现接口里的方法
/* define a method for circle (implementation of Shape.area())*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

// 面向接口编程，输入参数为Shape
/* define a method for shape */
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {

	//var circle1 Circle
	//circle1.radius=5
	circle := Circle{
		x:      0,
		y:      0,
		radius: 5,
	}
	rectangle := Rectangle{
		width:  10,
		height: 5,
	}

	fmt.Printf("Circle area: %f\n", getArea(&circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
