package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var numberOne, numberTwo int = 0, 0

	for {
		fmt.Print("Digite el numero 1: ")
		scanner.Scan()
		data, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			fmt.Println("Porfavor digite un numero valido ")
			continue
		}
		numberOne = int(data)
		break
	}

	for {
		fmt.Print("Digite el numero 2: ")
		scanner.Scan()
		data, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			fmt.Println("Porfavor digite un numero valido ")
			continue
		}
		numberTwo = int(data)
		break
	}

	fmt.Printf("La suma es: %d \n", Sum(numberOne, numberTwo))
}
