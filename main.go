package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		my_first_line := scanner.Text()
		my_first_word := strings.Fields(strings.ToLower(my_first_line))
		fmt.Println("Your command was:", my_first_word[0])
	}
}
