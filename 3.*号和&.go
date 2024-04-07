package main

import "fmt"

// * 当星号用在变量声明时，它表示该变量是一个指针类型，存储的是另一个变量的内存地址
func main() {
fmt.Println("hello world")
var x int =10
var ptr *int  = &x // ptr 是一个指向整数类型的指针，存储了变量 x 的地址
fmt.Println("打印效果",ptr,*ptr)
}

