package main

func tipo(i interface{}) string {
	switch i.(type) {
	case int, int64:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	case bool:
		return "booleano"
	case *int:
		return "ponteiro para inteiro"
	case *bool:
		return "ponteiro para booleano"
	case *string:
		return "ponteiro para string"
	default:
		return "não sei"
	}
}

func main() {

	println(tipo(2.3))
	println(tipo(1))
	println(tipo("Opa"))
	println(tipo(func() {}))
	println(tipo(true))
	println(tipo(&[]int{}))
	println(tipo(nil))

	var pi *int = nil
	println(tipo(pi))

	var ps *string = nil
	println(tipo(ps))

	teste := true
	var tb *bool = &teste
	println(tipo(tb))
}
