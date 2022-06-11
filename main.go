package main

import (
	"bufio"
	"os"
)

func main() {
	if input, err := os.Open("input.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandToParse := scanner.Text()
		}
	}
}
