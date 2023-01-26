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
만약 소문자로 시작하게 되면 private 함수로 인식하여 불러올 수 없다.
## 3. 모듈을 만들땐, 폴더를 만들고 그 안에 .go 파일을 만들어라
이후 해당 모듈 폴더에서 `go mod init [모듈경로; github.com/oiuyy207/learn-Go/01-25/myModule]`를 이용하여 `go.mod`파일을 만들어야한다.  
(추가사항) 해당모델은 로컬에서만 만들고, 배포를 아직 하지 않은 상태이기 때문에 추가작업이 필요하다.  

1. 로컬패키지 참조하도록 변경('go.mod' 안에 작성됨)
```
go mod edit -replace github.com/oiuyy207/learn-Go/01-25/myModule=../myModule
```
2. 패키지 설정
```
go mod tidy
```

