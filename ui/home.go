package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"temp-convert/utils"
)

var reader bufio.Reader = *bufio.NewReader(os.Stdin)

func Home() {
	println("Welcome to the Temperature Converter!")
	println("Please select an option:")
	println("1. Do conversion\n2. View conversion history\n0. Exit")
	println("Please enter your choice (1-3):")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		input()
	case 2:
		utils.ShowHistory()
		fmt.Print("Press Enter to go back...")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			Home()
		} else {
			os.Exit(1)
		}
	case 0:
		println("Exiting the program. Goodbye!")
		os.Exit(0)
	default:
		println("Invalid choice. Please select a valid option.")
		os.Exit(1)
	}
}
