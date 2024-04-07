package main

import "fmt"

func main() {
	var arr=[5]int{1,2,3,4,5}
	var slice=arr[0:3]
	var slice2=append(slice,6,7)

	fmt.Println("一个一个",slice2)
	var slice3=append(slice,slice2...)

	fmt.Println("一块一块的",slice3)


	// 数组拷贝
	var arr2 []int=[]int{1,2,3,4,5}
	var b=make([]int,10)
	copy(b,arr2)
	fmt.Println("b",b)
}