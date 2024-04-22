package utill

import "fmt"

//UpperCamelCase => public func
func TestFunction() {
	testFunction()
}

//lowerCamelCase => private func
func testFunction() {
	fmt.Println("testFunction")
}