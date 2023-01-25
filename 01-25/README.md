# 학습노트

## 1. Go는 컴파일 언어  
> 즉, 주축이 될 main.go가 필요!(단, 라이브러리를 만드는 등 컴파일이 필요없으면 만들지 않음)  
> Go 컴파일러는 `main.go` → `package main` → `func main(){}` 을 가장 먼저 찾고 실행한다!

(예시)**'main.go'**
```go
//주석은 지금처럼 '//'을 이용
package main 

import "fmt"

func main() {
    //파이썬과 달리 ''는 문자열로 인식 안됨!
    //문자열은 반드시 ""!
    fmt.Println("Hello, world!")
}
```
## 2. func을 export하고자 하면, func의 첫 글자는 반드시 대문자여야 한다.  
때문에 어떤 모듈을 불러오고 그 안의 함수를 불러올 때, 첫글자가 대문자이다.