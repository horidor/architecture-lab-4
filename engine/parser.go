package engine

import (
	"strings"
)

func Parse (toParse string) command {
	args := strings.Fields(toParse)



	cmd := args[0]
	cmdargs := args[1:]

	switch cmd {
	case "print":
		if len(args) < 2 {
			return &printCommand{"syntax error: not enough arguments"}
		}
		message := strings.Join(cmdargs, " ")
		return &printCommand{message}

	case "cat":
		if len(cmdargs) != 2 {
			return &printCommand{"syntax error: invalid arguments"}
		}

		return &catCommand{cmdargs[0], cmdargs[1]}
	
	case "stop":

		return &stopCommand{};
	default:
		return &printCommand{"syntax error: invalid command"}

	}
}