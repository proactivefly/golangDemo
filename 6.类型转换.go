package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 10
	var str =strconv.Itoa(num1)
	fmt.Printf("字符为：%T , s1=%q",str,str)

	// 字符换转为数字

	var sts="1234"
	var num2,_=strconv.Atoi(sts)
	fmt.Printf("类型为：%T,数值为%v",num2,num2)
}