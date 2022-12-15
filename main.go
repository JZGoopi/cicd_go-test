package main

import (
	"apple/calc"
	"fmt"
	"os"
)

func main() {
	var password = os.Getenv("password")
	fmt.Println(calc.Add(3,5), password)
}
