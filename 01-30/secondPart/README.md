# 학습노트

## 1. 포인터
포인터에 접근하려면 `&`를 앞에 붙일 것  
포인터에 있는 값을 보려면 `*`를 앞에 붙일 것
```go
a := 10
b := &a
fmt.Println(b) //0xc0000a6058 <- 뭐이런 종류의 값이 나옴
fmt.Println(*b) // 10(a에 저장된 값)

*b = 12 //b 포인터의 변수(a)를 수정
fmt.Println(a) // 12
```


## 2. arrays와 slices
```go
//arrays 선언 방법 [길이]타입{값,값,...}
arr := [3]int{1,2,3}
```
```go
//slices 선언방법 []타입{값,값,...}
slic := []int{1,2,3}
//slices에 값을 이어 추가하려면 append함수 사용
slic = append(slic,4)//append함수는 자체적으로 수정하지 않음
```

## 3. maps
```go
//map 선언 방법 map[키 타입]값 타입{키:값, 키:값,...}
ma = map[string]int{"A": 1, "B": 2}

//key, value쌍으로 가지고 있기 때문에 for문에서 쌍별 순서대로 가져올 수 있음
for k,v := range ma{
    //처리
}
```