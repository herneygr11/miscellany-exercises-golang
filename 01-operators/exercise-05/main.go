package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/herneygr11/exercise-05/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var base, height float64 = 0, 0

	for {
		fmt.Println("Ingresa la base:")
		scanner.Scan()
		data, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Debes ingresar un valor valido")
			continue
		}
		base = data
		break
	}

	for {
		fmt.Println("Ingresa la altura")
		scanner.Scan()
		data, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Debes ingresar un valor valido")
			continue
		}
		height = data
		break
	}

	fmt.Printf("El area del rectángulo es: %.2f\n", services.Area(base, height))

}
