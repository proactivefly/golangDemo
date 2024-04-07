package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random number between min and max
func randomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}


// 返回两个值的函数
func getData() (int, int, int) {
	// 返回3个值
	return 2, 4, 8
}
func main() {
	// Generate a random number between 1 and 10 and print it
	num := randomNumber(1, 10)
	fmt.Println("Random number:", num)
	_, b, _ := getData()
	fmt.Println("b:", b)
}

