package main

import (
	"fmt"
	"strings"
)

/*
*문단단위로 주석 처리하는 방법
 */

// 함수를 작성할때 변수마다 타입지정 + 결과 타입도 지정 필요
// 여러개가 같은 타입이면 아래와 같이 ','로 나열하고 하나만 적어도 됨
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// / ...은 여러 매개변수를 한번에 받는 것!
func manyArgument(str ...string) {
	//이 경우 여러개의 입력을 받아 자동적으로 array로 만들어 취급!
	fmt.Println(str)
}

func main() {
	const name1 string = "Jaeyong" //const는 값을 변경할 수 없는 변수! 즉 수학에서 상수처럼 취급!
	//var name2 string = "재용"        //var는 변수! 나중에 수정가능!
	name2 := "재용" //위의 var를 이용하는 것과 동일. 타입은 넣는 값을 보고 알아서 지정

	fmt.Println("const 변수 : ")
	fmt.Println(name1)

	fmt.Println("\nvar 변수 : ")
	fmt.Println(name2)
	name2 = "아로나"
	fmt.Println(name2)

	fmt.Println("------------------------------")
	//다음 챕터

	fmt.Println(multiply(2, 2))

	fmt.Println("------------------------------")
	//다음 챕터

	totalLength, upperName := lenAndUpper("JaeYong")
	fmt.Println(totalLength, upperName)
	manyArgument("재용", "jaeyong", "JAEYONG", "Jae", "Yong")
}
