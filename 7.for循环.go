package main

import "fmt"

// 定义一个名为Person的结构体，提高代码的可重用性和可读性
type Person struct {
	Name string
	Age  int
}

func main() {
	// 数组遍历，没有变化
	var arr = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i]) // 使用fmt.Println替代println，准备就绪处理潜在的错误
	}
}