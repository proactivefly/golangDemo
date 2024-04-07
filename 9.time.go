package main

import (
	"fmt"
	"time"
)
func main(){
	var now time.Time=time.Now()
	fmt.Println(now)
	// 注意区分大小写
	fmt.Printf("%v什么格式的呢:%T\n",time.Now(),time.Now())
	//这个参数字符串的各个数字必须是固定的，必须这样写

	/*
		DateTime   = "2006-01-02 15:04:05"
		DateOnly   = "2006-01-02"
		TimeOnly   = "15:04:05"
	*/
	fmt.Println("当前时间是：",time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("当前时间是：",time.Now().Format(time.DateTime))
	fmt.Println("当前年份是：",now.Year())
	fmt.Println("当前月份是：",int(now.Month()))
	fmt.Println("当前天数是:",now.Day())
	fmt.Printf("时：%v \n",now.Hour())
	fmt.Printf("分：%v \n",now.Minute())
	fmt.Printf("秒：%v \n",now.Second())
	
}