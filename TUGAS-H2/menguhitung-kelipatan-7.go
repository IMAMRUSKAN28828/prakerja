package main

import (
	"fmt"
)

func isMultipleOf7(n int) bool {
	return n%7 == 0
}

func main() {
	number := 21 // Ganti dengan bilangan yang ingin diuji

	if isMultipleOf7(number) {
		fmt.Printf("%d adalah kelipatan dari 7.\n", number)
	} else {
		fmt.Printf("%d bukan kelipatan dari 7.\n", number)
	}
}
