// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 1
	// print the value of the local variable "val"
	fmt.Println(val)
	// print the value of the package-level variable "val"
	printPackageVal()
	// update the package-level variable "val" and print it again
	updatePackageVal()
	printPackageVal()
	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val = %v\n", &val, &val)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 100
	fmt.Println("local val = ", val)
}

func printPackageVal() {
	fmt.Println(val)
}

func updatePackageVal() {
	val = "global2"
}