package engine

import (
	"strings"
)

func Parse (toParse string) Command {
	args := strings.Fields(toParse)

	if len(args) < 2 {
		return &PrintCommand{"syntax error: not enough arguments"}
	}


	cmd := args[0]
	cmdargs := args[1:]

	switch cmd {
	case "print":
		message := strings.Join(cmdargs, " ")
		return &PrintCommand{message}

	case "cat":
		if len(cmdargs) != 2 {
			return &PrintCommand{"syntax error: invalid arguments"}
		}

		return &CatCommand{cmdargs[0], cmdargs[1]}
	
	default:
		return &PrintCommand{"syntax error: invalid command"}

	}
}