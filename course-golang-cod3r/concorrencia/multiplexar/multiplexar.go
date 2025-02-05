package main

import (
	generatorHtml "coder_golang/concorrencia/generator"
	"fmt"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		generatorHtml.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		generatorHtml.Titulo("https://github.com/joaodiasdev", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
