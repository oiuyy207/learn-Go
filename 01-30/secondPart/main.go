package main

import "fmt"

func main() {
	a := 2
	b := a  //a의 값이 복사되서 들어감
	p := &a //a의 메모리 주소값(포인터)이 들어감
	a = 10  //b에 영향 없음

	fmt.Println(&a, &b)
	fmt.Println(&a, p)
	fmt.Println(*p)

	*p = 12 //'p에 들어있는 주소'의 값을 수정

	fmt.Println(a)

	fmt.Println("--------------")

	names := [5]string{"Alice", "Ben", "Catherine"}
	fmt.Println(names)
	//index순서는 0,1,2,3,...
	names[3] = "Daniel"
	names[4] = "Echo"
	fmt.Println(names)

	slic := []int{1, 2, 3}
	fmt.Println(slic)
	slic = append(slic, 4)
	fmt.Println(slic)

	fmt.Println("--------------")

	arona := map[string]string{"apperance": "VERY CUTE", "meme": "Mol?lu"}
	fmt.Println(arona)
	for key, value := range arona {
		fmt.Println(key, value)
	}
}
