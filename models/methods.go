package models

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Printf("\n>>>>>>> menu list\n")
	p := Head
	for p != nil {
		fmt.Printf("%s\t-\t%s\n", p.Cmd, p.Desc)
		p = p.Next
	}
	fmt.Printf(">>>>>>>\n\n")
}

func Quit() {
	fmt.Printf("\n>>>>>>> Bye!\n")
	os.Exit(0)
}
