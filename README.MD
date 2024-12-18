# Temperature Converter

A simple and efficient command-line temperature conversion tool written in Go.

## Overview

This program allows quick and accurate temperature conversions between Celsius, Fahrenheit, and Kelvin scales through an interactive command-line interface.

## Features

✨ **Multiple Scale Support**
- Celsius (°C)
- Fahrenheit (°F) 
- Kelvin (K)

🛡️ **Robust Error Handling**
- Validates numeric inputs
- Ensures correct scale specification
- Checks input format

🔧 **User-Friendly Interface**
- Interactive command prompt
- Clear input instructions
- Formatted output display

## Getting Started

### Prerequisites
- Go 1.23.4 or later ([Download Go](https://golang.org/dl/))

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Airsane/TemperatureConvertor.git
cd TemperatureConvertor
```

2. Build the program:

```bash
go build
```

3. Run the program:

```bash
./temperature
```

## Usage

1. Enter the temperature value followed by the scale (C, F, or K).
2. The program will display the converted temperature in all three scales.
3. Type "exit" to quit the program.

## Example

```bash
Enter temperature value: 25 C
25.00°C = 77.00°F = 298.15 K
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is open-sourced under the MIT License - see the LICENSE file for details.

