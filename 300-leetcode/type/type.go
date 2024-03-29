package _type

import (
	"fmt"
	"reflect"
)

func printType(whoami interface{}) {
	switch whoami.(type) {
	case int:
		fmt.Println("I am an int")
	case string:
		fmt.Println("I am a string")
	case bool:
		fmt.Println("I am a bool")
	default:
		// Если тип не известен
		fmt.Printf("Unknown type: %s\n", reflect.TypeOf(whoami))
	}
}

// добавить код, который выведет в терминал
// тип переданной переменной whoami

func main() {
	printType(42)
	printType("im string")
	printType(true)
}
