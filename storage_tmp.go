//usr/bin/env go run "$0" "$@"; exit "$?"
package main

import (
	"fmt"
)

func main() {
	i := "dog"
	switch i {
	case "hey":
		fmt.Println("one")
	case "dog":
		fmt.Println("two")
	case "cat":
		fmt.Println("three")
	}
}
