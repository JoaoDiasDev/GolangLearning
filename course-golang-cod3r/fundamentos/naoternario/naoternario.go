package main

import "fmt"

// Não tem ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	fmt.Println(obterResultado(6.2))
}
