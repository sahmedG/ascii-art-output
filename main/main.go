package main

import (
	"asciiART"
	"fmt"
	"os"
	"strings"
)

func IsFlagPassed(flag string) bool {
	for _, arg := range os.Args {
		if arg == flag {
			return true
		}
	}
	return false
}

func main() {
	//* parsing arguments
	var atCounter int

	if len(os.Args[1:]) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	for i := 0; i < len(os.Args); i++ {
		if IsFlagPassed("--output=") {
			continue
		} else {
			atCounter = i
			break
		}
	}
	args := os.Args[atCounter:]
	//* handling invalid arguments
	if len(args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	} else if len(args) >= 3 && len(args) <= 4 {
		//* handling null input text
		if os.Args[1] == "" {
			fmt.Print()
		} else {
			//* handling invalid charachters
			for _, v := range args[0] {
				if v < ' ' || v > '~' {
					fmt.Println("Invalid charachters detected")
					return
				}
			}
			//* executing function
			if strings.Contains(args[1], "--output") && len(args) > 3{
				outputfile := strings.ReplaceAll(args[1], "--output=", "")
				asciiART.Output(true, outputfile, args[2], (args[3]))
			}else if strings.Contains(args[1], "--output=") {
				outputfile := strings.ReplaceAll(args[1], "--output=", "")
				asciiART.Output(true, outputfile, args[2], "standard")
			} else {
				asciiART.Output(false, "", args[1], args[2])
			}
		}
	//} else if len(os.Args) == 3 {
	//	asciiART.Output(true, "", args[1], args[2])
	} else if len(os.Args) == 2 {
		asciiART.Output(false, "", args[1], "standard")
	} else if len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	}
}
