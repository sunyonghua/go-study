package main

import "fmt"

func main() {
	// 使用字面量创建 Map
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(m)

	// 获取值
	apple := m["apple"]
	fmt.Println(apple)
}
