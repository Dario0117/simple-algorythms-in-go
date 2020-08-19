package main

import (
	"os"

	"github.com/Dario0117/simple-algorythms-in-go/algorythms"
)

func main() {
	args := os.Args[1:]
	switch args[0] {
	case "clc":
		algorythms.RunCalculator(args[1:])
	case "rmn":
		algorythms.RunRomanNumeralParser(args[1:])
	}
}
