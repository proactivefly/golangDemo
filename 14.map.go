package main

import "fmt"
func main (){
	var mapp=make(map[string]string,10)

	//map可以存放10个键值对
	mapp["美利达"]="公爵600"
	fmt.Println("mapp",mapp)

	fmt.Printf("格式为：%T\n",mapp)


	// 方式二
	c:=map[int]string{
		1:"a",
		2:"b",
		3:"c",
	}

	fmt.Println("c",c)


	// delete
	delete(c,1)
	fmt.Println("c",c)

	// 查找
	value,flag:=c[2]
	fmt.Println(value)
	fmt.Println(flag)
}