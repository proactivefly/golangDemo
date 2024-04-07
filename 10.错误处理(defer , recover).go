package main

import "fmt"
func main(){
	defer func(){
		err:=recover()
		fmt.Println(err)
		fmt.Printf("err类型为%T \n",err)
		if err!=nil{
			fmt.Println("错误已经捕捉err:",err)
		}
	}()
	num1:=10
	num2:=0
	result:=num1/num2
	fmt.Println("结果是:",result)
}