package main

func main() {

	i := 0

	for i < 10 {
		i++
		println("A:", i)

	}

	for i := 0; i < 20; i += 2 {
		println("B:", i)
	}

	for i := 0; i < 10; i++ { // for clássico
		println("C:", i)
	}

	for {
		println("Loop infinito")
	}
}
