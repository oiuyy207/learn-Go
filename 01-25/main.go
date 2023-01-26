package main

import (
	"fmt"

	"github.com/oiuyy207/learn-Go/01-25/myModule"
	//go.mod에는 폴더까지로 init했으니, 그 안의 go파일까지로 경로지정!
	//만약 모듈을 만들고 자동완성 쓰려면 먼저 init을 하고 기능 사용하자!
)

func main() {
	fmt.Println("Hello, world!")
	myModule.SayHello()
}
