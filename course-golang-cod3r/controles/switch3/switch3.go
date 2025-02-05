package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float64, float32:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	case bool:
		return "boll"
	default:
		return "nao sei"
	}

}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("ola"))
	fmt.Println(tipo(true))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))

}
