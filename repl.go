package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startfunc() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		my_first_line := scanner.Text()
		my_first_word := cleanInput(my_first_line)
		fmt.Println("Your command was:", my_first_word[0])
	}
}

func cleanInput(text string) []string {
	var my_string []string

	my_text := strings.ToLower(text)

	my_text = strings.TrimSpace(my_text)

	my_string = strings.Fields(my_text)

	return my_string
}
