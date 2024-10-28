package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите целое число: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("Вы ввели число: %v\n", input)
}
