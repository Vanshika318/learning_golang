package main

import "fmt"

func main() {
	slice_int := []int{80, 99, 10, 15, 20, 30, 40, 33, 9, 47}

	for _, n := range slice_int {
		if n%2 == 0 {
			fmt.Println(n, "is an even number")
		} else {
			fmt.Println(n, "is an odd number")
		}
	}
}
