# 학습노트

## 1. 변수 생성
Go에서 변수를 선언할 때 타입도 같이 지정한다.  
즉, 코드 중간에 **변수에 대해 타입변경이 파이썬처럼 자유롭게 되지는 않는다!**

`const`는 변경할 수 없는 변수  
```go
const Value int = 123
Value = 234 //불가능! 이 변수에 넣을 수 없다고 오류 출력
``` 
`var`은 변경 가능한 변수. 일반적으로 많이 사용.  
축약형 '`:=`'은 `func`안에서만 동작한다. 만약 `func`바깥에서 변수를 지정하려면 `var [변수명] [타입] = [값]`을 사용해야함.
```go
//var Value int = 123
Value := 123 
Value = 234 //쌉가능
```

## 2. 함수 작성
어떤 값을 `return`하는 함수를 작성하기 위해서는 넣는 매개변수와 결과 모두 타입지정을 해주어야 한다.
```go
func Example(a int, b int) int{
    return a+b
}
```

### 3. Go 함수의 장점 : 여러개의 `return`
Go 함수는 하나의 함수에 여러개의 `return`을 지정할 수 있다.  
여러개의 값을 내놓는다면, `_`을 이용해서 값 무시를 할 수 있다.(파이썬 처럼)
```go 
func Example(name string) (int, string){
    return len(name), strings.ToUpper(name)
}
```

### 4. 여러 매개변수 받기
같은 타입의 여러개의 매개변수를 받으려면 타입지정할 때 `...`을 앞에 써주면 된다.
```go
func manyArgument(str ...string) {
	//이 경우 여러개의 입력을 받아 자동적으로 array로 만들어 취급!
	fmt.Println(str)
}
```