package main

import (
	"bufio"
	"os"

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
	 		eventLoop.Post(engine.Parse(commandToParse))
	 	}
	}

	eventLoop.AwaitFinish()
}
