package main

import (
	"fmt"

	//go.mod에는 폴더까지로 init했으니, 그 안의 go파일까지로 경로지정!
	//만약 모듈을 만들고 자동완성 쓰려면 먼저 init을 하고 기능 사용하자!
	mm "github.com/oiuyy207/learn-Go/01-25/myModule/myModule" // python의 as와 같이 앞에서 축약어 지정 가능!
)

func main() {
	fmt.Println("Hello, world!")
	mm.SayHello()

}
