package main

import (
	"fmt"
	"strings"
)

// return할 값을 함수 선언 단에서 '생성' 가능
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I`m done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) (result int) {
	result = 0
	for index, number := range numbers {
		fmt.Println("index, value :", index, ",", number)
		result += number
	}
	return
}
func main() {
	length, uppercase := lenAndUpper("jaeyong")
	fmt.Println(length, uppercase)

	fmt.Println("---------------------------")
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
