package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	triangle := Triangle{base: 0, height: 0}

	fmt.Println("Bienvenido: ")

	for {
		fmt.Println("\nDigite la base")
		scanner.Scan()
		base, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error: Debe digitar un valor valido.")
			continue
		}
		triangle.base = base
		break
	}

	for {
		fmt.Println("\nDigite la altura")
		scanner.Scan()
		height, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error: Debe digitar un valor valido.")
			continue
		}
		triangle.height = height
		break
	}

	area := triangle.Area()

	fmt.Printf("\nEl area del triangulo es: %.2f \n", area)
}
