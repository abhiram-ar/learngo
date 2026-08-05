package main

import "fmt"

func main() {
	slice := make([]int, 4)
	fmt.Printf("%p\n", &slice)

	slice = append(slice, 2)
	fmt.Printf("%p\n", slice)

	slice = append(slice, 3)
	fmt.Printf("%p\n", slice)

	slice = append(slice, 4)
	fmt.Printf("%p\n", slice)

	slice = append(slice, 4)
	fmt.Printf("%p\n", slice)

	slice = append(slice, 4)
	fmt.Printf("%p\n", slice)

	slice = append(slice, 4)
	fmt.Printf("%p\n", slice)

}
