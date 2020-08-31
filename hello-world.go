package main

import (
	"fmt"
	"reflect"
	"runtime"
)

/*
var (
	name   string  = "Nigel"
	course string  = "Docker deep dive"
	module float64 = 3.2
)
*/

func main() {

	name := "Nigel"
	course := "Docker Deep Dive"
	module := 3.2
	ptr := &module

	fmt.Println("Hello from", runtime.GOOS)

	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Course is set to", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memory address of *module* variable is", ptr,
		"and the value of *module is", *ptr)

}
