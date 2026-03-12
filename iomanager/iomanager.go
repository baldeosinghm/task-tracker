package iomanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func WelcomeMsg() {
	fmt.Println("Hi and welcome to Task Tracker!")
	fmt.Print("Your command: ")
}

func InputParser() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, errors.New("Failed to parse user input.")
	}
	line := scanner.Text()
	words := strings.Fields(line)

	return words, nil
}
