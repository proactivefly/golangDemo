package main

import "fmt"


type Person struct {
	Name string
	Age  int
	Sex  string
}
func main (){


	var xiaohong Person= Person{
		Name:"小红",
		Age:18,
		Sex:"女",
	}

	// var arr=[2]int{1,2}
	// var arr1 []int
	// fmt.Println(arr)
	// fmt.Println(arr1)
	fmt.Println(xiaohong)
	// 创建结构体实例
	xiaoming:= Person{
		Name:"小明",
		Age:18,
		Sex:"男",
	}
	fmt.Println(xiaoming)


	fmt.Println("————————————————-")

	var point *Person=new(Person)
	fmt.Println("point",point)
	point.Name="冯晓航"
	point.Age=18
	point.Sex="男"
	fmt.Println(*point)

	fmt.Println("————————————————-")


	var point1 *Person=&Person{
		Name:"冯晓航",
		Age:18,
		Sex:"男",
	}
	fmt.Println(*point1)
	
}