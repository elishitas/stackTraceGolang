package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	doSomething()
}

func doSomething() {
	// Generar un error
	err := someErrorFunc()
	if err != nil {
		fmt.Println("Ocurrió un error:")
		fmt.Println(err)
		debug.PrintStack()
	}
}

func someErrorFunc() error {
	return fmt.Errorf("Este es un error generado")
}

//El stack trace muestra la secuencia de llamadas de función
// que condujeron al error.
