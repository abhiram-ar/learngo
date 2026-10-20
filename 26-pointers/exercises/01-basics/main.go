// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import (
	"fmt"
)

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var asusPtr *computer

	// compare the pointer variable to a nil value, and say it's nil
	if asusPtr == nil {
		fmt.Println("asusPtr is nil")
	}

	// create an apple computer by putting its address to a pointer variable
	appleComputerPtr := &computer{"apple"}

	// put the apple into a new pointer variable
	appleComputerPtr2 := appleComputerPtr

	// compare the apples: if they are equal say so and print their addresses
	if appleComputerPtr == appleComputerPtr2 {
		fmt.Printf("Apples are equal %p  %p\n", appleComputerPtr, &appleComputerPtr2)
	}

	// create a sony computer by putting its address to a new pointer variable
	sonyComputerPtr := &computer{"sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if appleComputerPtr != sonyComputerPtr {
		fmt.Printf("Apple and sony are not equal %p %p\n", appleComputerPtr, sonyComputerPtr)
	}

	// put apple's value into a new ordinary variable
	appleComputerVal := *appleComputerPtr

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("appleComputerPtrAddress=%p appleComputerPtr=%p appleComputerAddress=%p\n", &appleComputerPtr, appleComputerPtr, &appleComputerVal)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *appleComputerPtr == appleComputerVal {
		// print the values:
		// the value that is pointed by the apple and the new variable
		fmt.Printf("appleComputerPtr is equal to new var: %+v\n", appleComputerVal)
	}

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sonyComputerPtr, "hp")
	fmt.Println("sony's new brand", sonyComputerPtr.brand)

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	value := getValue(appleComputerPtr)
	fmt.Printf("apple value %+v\n", value)

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("New dell computer      : add=%p\n", newComputer("dell"))
	fmt.Printf("New lenovo computer    : add=%p\n", newComputer("lenovo"))
	fmt.Printf("New samsunng computer  : add=%p\n", newComputer("samsung"))

}

func change(computer *computer, newBrand string) {
	computer.brand = newBrand
}

func getValue(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
