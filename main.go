package main

import "fmt"

func main() {
	length_convert()
}

func temp_convert() {
	fmt.Print("Enter Fahrenheit Value: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celcius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Celcius value: %.2fC ", celcius)
}

func length_convert() {
	fmt.Print("Enter length in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meter := feet * 0.3048
	fmt.Printf("Length in meters: %.2fm ", meter)
}
