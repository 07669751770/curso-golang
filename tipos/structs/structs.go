package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name        string
	dateOfBirth time.Time
}

func (p Person) getAge() uint {
	hoje := time.Now()

	// Calcula a diferença de anos
	idade := hoje.Year() - p.dateOfBirth.Year()

	// Verifica se ainda não fez aniversário este ano
	if hoje.YearDay() < p.dateOfBirth.YearDay() {
		idade--
	}
	return uint(idade)
}

func main() {
	p := Person{
		Name:        "Bruno Pueyo",
		dateOfBirth: time.Date(1977, time.March, 4, 0, 0, 0, 0, time.UTC),
	}
	fmt.Printf("Nome: %s\nIdade: %d\n", p.Name, p.getAge())
}
