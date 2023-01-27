# 학습노트

## 1. 함수에서 `return`할 변수를 함수 선언단에서 생성할 수 있다
주의! `return`할 변수를 '생성'한 것이기 때문에 다시 생성할 필요가 없음.  
또한, 이미 `return`할 값을 지정하였기 때문에 `return`만 적어넣어도 충분함(*naked return*).  
```go
func Example(a int) result int {
    result = a*2 //생성 하지 X
    return
}
```

## 2. `defer`기능
`defer`은 함수 내에서 `return`이후에 추가 동작을 지시하는 것!
```go
func Example(a int) (result int) {
    defer fmt.Println("I`m done!")
    result = a+10
    return
}
```
이 경우, 매개변수 `a`에 10을 더하고 "I`m done!"을 출력할 것이다.

## 3. Go에서 반복문 구문은 `for`하나만 존재!
`i`를 반복문 내에서 지정하고 하나씩 사용하는 방법을 파이썬과 비교해보면,
```python
#파이썬에서의 방법
for i in range(6):
    print(i)
```
```go
//Go 에서의 방법
for i:=0; i<6; i++{
    fmt.Println(i)
}
```

파이썬처럼 한 array에서 하나씩 받아서 print하는 코드를 작성하려면,
```python
#파이썬에서의 방법
array1 = [1,2,3,4,5,6]
for i in array1:
    print(i)
```
```go
//Go에서의 방법
array1 := [1,2,3,4,5,6]
//range에서 받는 변수가 한개면 값이 아닌 인덱스를 준다!
//즉, 아래와 같이 [인덱스 받을 변수],[값을 받을 변수]로 설정해줘야함.
for _, i := range array1{
    fmt.Println(i)
}
```