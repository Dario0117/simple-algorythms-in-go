package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	var result float64
	for {
		var selectedOption uint8

		if len(args) != 0 {
			if optionFromArgs, err := strconv.ParseInt(args[0], 10, 8); err == nil && optionFromArgs > 0 && optionFromArgs < 6 {
				selectedOption = uint8(optionFromArgs)
			} else {
				fmt.Println("❌ Please, select a valid option")
				return
			}
		} else {
			selectedOption = printMainMenu()
		}

		switch selectedOption {
		case 1:
			result = float64(printAddMenu())
		case 2:
			result = float64(printSubtractMenu())
		case 3:
			result = float64(printMiltiplyMenu())
		case 4:
			result = float64(printDivideMenu())
		case 5:
			fmt.Println("See ya!")
			return
		}
		fmt.Printf("🎊 The result is: %f 🎊\n", result)
	}
}

func printLineInsideMenu(text string) {
	fmt.Printf("# %-51s #\n", text)
}

func prompt(message string) string {
	var input string
	fmt.Printf(message + "> ")
	fmt.Scanln(&input)
	return input
}

func printMainMenu() uint8 {
	fmt.Println(strings.Repeat("#", 55))
	fmt.Println("# Welcome to this simple but functional calculator 🤡 #")
	printLineInsideMenu(strings.Repeat("-", 51))
	printLineInsideMenu("What you wanna do?")
	printLineInsideMenu("1 - For add")
	printLineInsideMenu("2 - For subtract")
	printLineInsideMenu("3 - For multiply")
	printLineInsideMenu("4 - For divide")
	printLineInsideMenu("5 - Exit")
	fmt.Println(strings.Repeat("#", 55))
	for {
		input := prompt("")
		if selectedOption, err := strconv.ParseInt(input, 10, 8); err == nil && selectedOption > 0 && selectedOption < 6 {
			return uint8(selectedOption)
		}
		fmt.Println("❌ Please, select a valid option")
	}
}

func readValidNumber() float64 {
	for {
		input := prompt("Please enter a number ")
		if number, err := strconv.ParseFloat(input, 10); err == nil {
			return number
		}
		fmt.Println("❌ Please, insert a valid number")
	}
}

func printAddMenu() float64 {
	fmt.Println(strings.Repeat("#", 55))
	printLineInsideMenu(strings.Repeat(" ", 12) + ">>>>> Add two numbers <<<<<")
	fmt.Println(strings.Repeat("#", 55))
	input1 := readValidNumber()
	input2 := readValidNumber()
	return input1 + input2
}

func printSubtractMenu() float64 {
	fmt.Println(strings.Repeat("#", 55))
	printLineInsideMenu(strings.Repeat(" ", 9) + ">>>>> Subtract two numbers <<<<<")
	fmt.Println(strings.Repeat("#", 55))
	input1 := readValidNumber()
	input2 := readValidNumber()
	return input1 - input2
}

func printMiltiplyMenu() float64 {
	fmt.Println(strings.Repeat("#", 55))
	printLineInsideMenu(strings.Repeat(" ", 9) + ">>>>> Miltiply two numbers <<<<<")
	fmt.Println(strings.Repeat("#", 55))
	input1 := readValidNumber()
	input2 := readValidNumber()
	return input1 * input2
}

func printDivideMenu() float64 {
	fmt.Println(strings.Repeat("#", 55))
	printLineInsideMenu(strings.Repeat(" ", 9) + ">>>>> Divide two numbers <<<<<")
	fmt.Println(strings.Repeat("#", 55))
	input1 := readValidNumber()
	var input2 float64
	for input2 == 0 {
		input2 = readValidNumber()
		if input2 == 0 {
			fmt.Println("❌ We can't divide by zero")
		}
	}
	return input1 / input2
}
