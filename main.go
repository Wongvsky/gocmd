package main

import (
	"bufio"
	"fmt"
	"menu/models"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(">>>>>>> Hello! Please input a command.")
	var cmd string
	for {
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		p := models.Head
		for p != nil {
			if cmd == p.Cmd {
				fmt.Printf(">>>>>>> %s\n", p.Desc)
				if p.Handler != nil {
					p.Handler()
				}
				break
			}
			p = p.Next
		}
		if p == nil {
			fmt.Println(">>>>>>> This is a wrong cmd!")
		}
	}
}
