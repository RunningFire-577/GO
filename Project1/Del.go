package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := Del(a1, 2)
	//a1 = append(a1[:1], a1[1+1:]...)
	//for _, y := range a1 {
	//	println(y)
	//}
	fmt.Print(a2)
	a3 := []string{"你好啊", "扑街", "衰仔"}
	a4 := Del(a3, 1)
	fmt.Print(a4)

}
