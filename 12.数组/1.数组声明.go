package main

import "fmt"

func main(){
	fmt.Println("hello world")
	var arr=[3]string{"x","y","z"}
	fmt.Println(arr)


	for key,value :=range arr{
		fmt.Printf("索引key:%v,值为：%v \n",key,value)
	}

	// 这里的三个点 ... 表示数组的长度由初始化列表中的元素数量决定，即根据后续花括号 {} 中提供的元素自动推断数组长度
	var arr2 =[...]int{1,2,3,4,5}
	fmt.Println(arr2)

	// 长度属于类型的一部分
	fmt.Printf("arr2类型为%T \n",arr2)
	
	var arr3 [3]int
	arr3[0]=1
	arr3[1]=2
	arr3[2]=3
	fmt.Println(arr3)



	arr4:=[3]int{1,2,3}
	fmt.Println(arr4)
}