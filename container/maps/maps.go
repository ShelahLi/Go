package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "lth",
		"course":  "golang",
		"site":    "bupt",
		"quality": "notbad",
	}

	// 空的map
	m2 := make(map[string]int) // m2 == empty map

	// nil的map
	var m3 map[string]int // m3 == nil

	fmt.Println("m, m2, m3:")
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map m")
	//遍历map, key在map里面是无序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(`m["course"] =`, courseName)

	//如果key值输入错误怎么办？会输出空串
	//判断key存不存在， ok
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}



	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)
	// 删除map里面的元素
	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
	fmt.Print(m)
}
