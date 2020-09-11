# Simple algorythms in go
This don't have a special purpose, just practicing Go

## Programs inside this  repo

### Calculator (menu and cli params)

#### Instructions

- For menu version run : `go run main.go clc`
- For cli version run: `go run main.go clc [option]`

Options: 
```
1 - For add
2 - For subtract
3 - For multiply
4 - For divide
5 - For Arithmetic expression
6 - Exit
```

Example:
```bash
» go run main.go clc 5
#######################################################
#            >>>>> Arithmetic expression <<<<<        #
#######################################################
Insert a expression > (53+5/(4+3))
🎊 The result is: 53.714286 🎊
```

- [x] Basic operations with two numbers supported individualy (+,-,*,/)
- [x] Arithmetic expression in one line supported (53+5/(4+3)) -> 53+5/7 -> 53+0.714286 -> 53.714286)


### Roman Numerals (menu and cli params)

#### Instructions

Note: Only IVXLCDM characters are valid

- For menu version run : `go run main.go rmn`
- For cli version run: `go run main.go rmn [Roman numeral]`

**Note: This converter doesn't really works 😂 (I need to check all the rules for the convertion so...), i just pushed this code because of the setup i made.**
