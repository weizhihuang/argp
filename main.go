// go:build !windows

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	args := "nil"
	if len(os.Args) > 1 {
		args = strings.Join(os.Args[1:], "\" \"")
		if len(os.Args) > 2 {
			args = "\"" + args + "\""
		}
		clipboard.WriteAll(args)
	}
	fmt.Println(args)
}
