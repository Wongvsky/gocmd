package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input the command. Type exit to quit.")
	var cmd string
	for cmd != "exit" {
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		switch cmd {
		case "exit":
		case "":
		default:
			fmt.Println("still developing......")
		}
	}
	fmt.Println("bye")
}
