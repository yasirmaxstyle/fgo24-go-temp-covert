package ui

import (
	"fmt"
	"os"
	"temp-convert/utils"
)

var choice int

func input() {
	fmt.Println("What do you want to convert?\n1. Celsius to Kelvin\n2. Celsius to Fahrenheit\n3. Celsius to Reaumur\n0. Back to main menu")
	fmt.Print("Please enter your choice (1-3): ")
	fmt.Scan(&choice)
	utils.Clear()
	switch choice {
	case 1:
		var celsius float64
		fmt.Print("Enter temperature in Celsius:")
		fmt.Scanf("%f", &celsius)
		fmt.Printf("Temperature in Kelvin: %.2f\n", utils.CtoK(celsius))
		utils.AddHistory(fmt.Sprintf("Celsius to Kelvin: %.2f", utils.CtoK(celsius)))
	case 2:
		var celsius float64
		fmt.Print("Enter temperature in Celsius:")
		fmt.Scanf("%f", &celsius)
		fmt.Printf("Temperature in Fahrenheit: %.2f\n", utils.CtoF(celsius))
		utils.AddHistory(fmt.Sprintf("Celsius to Fahrenheit: %.2f", utils.CtoF(celsius)))
	case 3:
		var celsius float64
		fmt.Print("Enter temperature in Celsius:")
		fmt.Scanf("%f", &celsius)
		fmt.Printf("Temperature in Reaumur: %.2f\n", utils.CtoR(celsius))
		utils.AddHistory(fmt.Sprintf("Celsius to Reaumur: %.2f", utils.CtoR(celsius)))
	case 0:
		Home()
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
		input()
	}
	fmt.Println("Do you want to do another conversion? (yes/no)")
	var response string
	fmt.Scan(&response)
	utils.Clear()
	if response == "yes" || response == "y" {
		input()
	} else if response == "no" || response == "n" {
		Home()
	} else {
		fmt.Println("Invalid response.")
		os.Exit(1)
	}
}
