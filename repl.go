package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var my_string []string

	my_text := strings.ToLower(text)

	my_text = strings.TrimSpace(my_text)

	my_string = strings.Fields(my_text)

	return my_string
}
