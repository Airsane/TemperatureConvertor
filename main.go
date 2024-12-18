package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Temperature struct {
	value float64
	scale string
}

func (t Temperature) toCelsius() float64 {
	switch t.scale {
	case "C":
		return t.value
	case "F":
		return (t.value - 32) * 5 / 9
	case "K":
		return t.value - 273.15
	default:
		return 0
	}
}

func (t Temperature) toFahrenheit() float64 {
	celsius := t.toCelsius()
	return celsius*9/5 + 32
}

func (t Temperature) toKelvin() float64 {
	celsius := t.toCelsius()
	return celsius + 273.15
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTemperature Converter")
		fmt.Println("Enter a value and scale (e.g., '32 F') or 'q' to quit:")

		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "q" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Error: Please enter both a value and a scale (C, F, or K)")
			continue
		}

		value, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Error: Invalid number")
			continue
		}

		scale := strings.ToUpper(parts[1])
		if !strings.Contains("CFK", scale) {
			fmt.Println("Error: Scale must be C, F, or K")
			continue
		}

		temp := Temperature{value: value, scale: scale}

		switch scale {
		case "C":
			fmt.Printf("%.2f°C = %.2f°F = %.2f K\n", value, temp.toFahrenheit(), temp.toKelvin())
		case "F":
			fmt.Printf("%.2f°F = %.2f°C = %.2f K\n", value, temp.toCelsius(), temp.toKelvin())
		case "K":
			fmt.Printf("%.2f K = %.2f°C = %.2f°F\n", value, temp.toCelsius(), temp.toFahrenheit())
		}
	}

	fmt.Println("Goodbye!")
}
