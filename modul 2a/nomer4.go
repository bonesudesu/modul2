package main

import "fmt"

func celciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}

func fahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func celciusToReamur(celcius float64) float64 {
	return celcius * 4 / 5
}

func celciusToKelvin(celcius float64) float64 {
	return celcius + 273.15
}

func main() {
	var celcius float64
	fmt.Print("Temperature in Celcius: ")
	fmt.Scanln(&celcius)

	fahrenheit := celciusToFahrenheit(celcius)
	reamur := celciusToReamur(celcius)
	kelvin := celciusToKelvin(celcius)

	fmt.Printf("Temperature in Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Temperature in Reamur: %.2f\n", reamur)
	fmt.Printf("Temperature in Kelvin: %.2f\n", kelvin)
}
