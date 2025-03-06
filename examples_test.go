package hutils_test

import (
	"fmt"

	template "atomicgo.dev/hutils"
)

func Example_demo() {
	fmt.Println(template.HelloWorld())
	// Output: Hello, World!
}

func ExampleHelloWorld() {
	fmt.Println(template.HelloWorld())
	// Output: Hello, World!
}
