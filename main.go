// nwe.go

package main

import "fmt"

func nwe(f int) int {
	if f == 1 {
		fmt.Print("num : ")
	} else if f == 0 {
		fmt.Println("error")
	}
	return f
}

func main() {
	var a int
	fmt.Print("test 1 ")
	fmt.Scan(&a)
	a = nwe(a)
}
