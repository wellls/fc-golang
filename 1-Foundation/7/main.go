package main

import "fmt"

func main() {
	salarios := map[string]int{"José": 1000, "João": 2000, "Maria": 3000}
	delete(salarios, "José")
	salarios["Jos"] = 5000

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["José"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
