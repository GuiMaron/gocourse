package main

import "fmt"

func main () {

	numbers := make([]int, 11)

	for i := range numbers {
		numbers[i] = i
	}

	for _, n := range numbers {

		if (n % 2) == 0 {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}

	}

}
