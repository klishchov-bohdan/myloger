package logger

import (
	"fmt"
	"log"
)

func PrintHello(name string) {
	fmt.Println("Hello, ", name)
}

func Fatal(i interface{}) {
	log.Fatal(i)
}
