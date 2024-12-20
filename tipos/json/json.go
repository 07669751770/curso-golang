package main

import (
	"encoding/json"
	"fmt"
)

type Documento struct {
	Numero uint   `json:"numero"`
	Tipo   string `json:"tipo"`
}

type Person struct {
	Name       string      `json:"name"`
	Age        int         `json:"age"`
	Documentos []Documento `json:"documentos"`
}

func main() {
	p := Person{Name: "Bruno Pueyo", Age: 47, Documentos: []Documento{{Numero: 7669751770, Tipo: "CPF"}, {Numero: 7669751770, Tipo: "RG"}}}
	jsonData, err := json.Marshal(p)
	printStringFromByteArray(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", p2)

	var p3 interface{}
	err = json.Unmarshal(jsonData, &p3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", p3)

}

func printStringFromByteArray(b []uint8) {
	for _, v := range b {
		fmt.Print(string(v))
	}
	fmt.Println()
}
