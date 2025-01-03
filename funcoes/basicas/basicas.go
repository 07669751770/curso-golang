package main

import "fmt"

func main() {
	f1()
	f2("Param1", "Param2")
	r3 := f3()
	fmt.Println("F3:", r3)
	r4 := f4("Param1", "Param2")
	fmt.Println("F4:", r4)
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println("F2: ", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1 string, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno1", "Retorno2"
}
