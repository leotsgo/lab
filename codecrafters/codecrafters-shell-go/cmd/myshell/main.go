package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the command: ", err)
		}

		input = strings.TrimSpace(input)
		command := strings.Split(input, " ")

		switch command[0] {
		case "exit":
			os.Exit(0)
			break
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(command[1:], " "))
			break
		case "type":
			switch command[1] {
			case "exit", "echo", "type", "cat":
				message := fmt.Sprintf("%s is a shell builtin\n", command[1])
				pathString := os.Getenv("PATH")
				path := checkFileInPath(pathString, command[0])

				if path != "" {
					message = fmt.Sprintf("%s is %s\n", command[0], path)
				}

				fmt.Fprintf(os.Stdout, "%s", message)
			default:
				fmt.Fprintf(os.Stdout, "%s: not found\n", command[1])
			}
			break
		default:

			fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
			break
		}
	}
}
