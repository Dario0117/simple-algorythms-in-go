package main

import (
	"bufio"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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
			if optionFromArgs, err := strconv.ParseInt(args[0], 10, 8); err == nil && optionFromArgs > 0 && optionFromArgs < 7 {
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
			result = float64(printArithmeticExpressionMenu())
		case 6:
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
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
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
	printLineInsideMenu("5 - For Arithmetic expression")
	printLineInsideMenu("6 - Exit")
	fmt.Println(strings.Repeat("#", 55))
	for {
		input := prompt("")
		if selectedOption, err := strconv.ParseInt(input, 10, 8); err == nil && selectedOption > 0 && selectedOption < 7 {
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

func printArithmeticExpressionMenu() float64 {
	fmt.Println(strings.Repeat("#", 55))
	printLineInsideMenu(strings.Repeat(" ", 11) + ">>>>> Arithmetic expression <<<<<")
	fmt.Println(strings.Repeat("#", 55))
	for {
		input1 := prompt("Insert a expression ")
		tr, err := parser.ParseExpr(input1)

		// Print AST
		// fset := token.NewFileSet()
		// ast.Print(fset, tr)

		if err != nil {
			fmt.Println("❌ Please, insert a valid expression")
		} else {
			result, err := eval(tr)
			if err != nil {
				fmt.Println("❌", err.Error())
			} else {
				return result
			}
		}
	}
}

func eval(exp ast.Expr) (float64, error) {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		result, err := evalBinaryExpr(exp)
		if err != nil {
			return 0, err
		}
		return result, nil
	case *ast.ParenExpr:
		return eval(exp.X)
	case *ast.Ident:
		return 0, errors.New("We can't operate with non-numeric values")
	case *ast.UnaryExpr:
		if exp.Op == token.SUB {
			val, _ := eval(exp.X)
			return -1 * val, nil
		}
		return eval(exp.X)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return float64(i), nil
		case token.FLOAT:
			i, _ := strconv.ParseFloat(exp.Value, 10)
			return float64(i), nil
		default:
			return 0, errors.New("Please, insert a valid expression")
		}
	}
	return 0, nil
}

func evalBinaryExpr(exp *ast.BinaryExpr) (float64, error) {
	left, err := eval(exp.X)
	if err != nil {
		return 0, err
	}
	right, err := eval(exp.Y)
	if err != nil {
		return 0, err
	}
	switch exp.Op {
	case token.ADD:
		return left + right, nil
	case token.SUB:
		return left - right, nil
	case token.MUL:
		return left * right, nil
	case token.QUO:
		if right == 0 {
			return 0, errors.New("We can't divide by Zero")
		}
		return left / right, nil
	}
	return 0, nil
}
