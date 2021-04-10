package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numberOne float64 = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese el numero: ")

		scanner.Scan()
		data, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Digite un numero valido")
			continue
		}

		numberOne = data
		break
	}

	fmt.Printf("El número elevado: %.2f\n", Squaring(numberOne))
}
