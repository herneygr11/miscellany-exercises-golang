package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/herneygr11/exercise-04/services"
)

var euros float64 = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese la cantidad de euros: ")

		scanner.Scan()
		data, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Digite una cantidad valida")
			continue
		}

		euros = data
		break
	}

	fmt.Printf("La cantida de euros es: %f y en dolares es: %f\n", euros, services.Converter(euros))
}
