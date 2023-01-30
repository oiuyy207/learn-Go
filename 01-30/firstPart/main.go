package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	//fmt.Println(koreanAge)
	return true

}

func switchTest(butten int) {
	switch flag := butten + 1; flag {
	case 1:
		fmt.Println("1번!")
	case 2:
		fmt.Println("2번!")
	case 3:
		fmt.Println("3번!")
	}
}

func canIDrink2(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18 && koreanAge < 65:
		return true
	case koreanAge >= 65:
		fmt.Println("주의 요망!")
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
	fmt.Println("---------------------")
	fmt.Println(canIDrink2(70))
	switchTest(2)

}
