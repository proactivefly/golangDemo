package main

import (
	"errors"
	"fmt"
)


func main(){
	var err=getData1()
	if err != nil {
		fmt.Println("你怎么这样啊",err)
	}
	
}

func getData1() error {
	// 模拟错误
	err := errors.New("模拟错误")
	return err
}