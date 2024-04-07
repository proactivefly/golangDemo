package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var slice = arr[1:3]
	fmt.Println(slice)
	fmt.Println("切片容量是",cap(slice))

	// 定义切片 ,切片类型，切片长度，切片容量
	qp:=make([]int,3,5)
	fmt.Printf("切片的类型是 %T\n",qp) // []int
	fmt.Println("容量是",cap(qp))


	// 遍历切片
	for i:=0 ;i<len(qp) ;i++ {
		fmt.Printf("key是%v---value是%v\n",i,qp[i])
	}
}