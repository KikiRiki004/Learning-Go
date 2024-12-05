package main

import (
	"fmt"
)

var f int = 2
var g = 3 //declaring with var works outside and inside the function
// var f, g int = 2, 3  //can be written in line

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a string //zero-value: ""
	var b int    //zero-value: 0
	var c bool   //zero-value: false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// var student3 string
	// student3 = "Johana"
	// fmt.Println(student3)

	d := 1 //can be used only inside the function
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(g)

	// var a, b = 6, "Hello" //a=6, b=hello
	// c, d := 7, "World!"   //c=7, b=world

	//myVariableName = "John"   //camel case
	//MyVariableName = "John"   //pascal case
	//my_variable_name = "John" //snake case

}
