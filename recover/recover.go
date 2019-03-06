package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()
	panic(1213)		//不认识的错误 还是panic

	b := 0
	a := 5/b
	fmt.Println(a)
}
func main() {
	tryRecover()
}
