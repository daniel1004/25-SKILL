package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите информацию: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("Вы ввели : %v\n", input)

}
