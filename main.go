package main

import (
	"bufio"
	"os"
	"fmt"

	"github.com/horidor/architecture-lab-4/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open("input.txt"); err == nil {
	 	defer input.Close()
	 	scanner := bufio.NewScanner(input)
	 	for scanner.Scan() {
	 		commandToParse := scanner.Text()
	 		if err := eventLoop.Post(engine.Parse(commandToParse)); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
			}
	 	}
	}

	eventLoop.AwaitFinish()

	if input, err := os.Open("test.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandToParse := scanner.Text()
			if err := eventLoop.Post(engine.Parse(commandToParse)); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
			}
		}
	}

}
