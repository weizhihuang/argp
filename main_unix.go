//go:build !windows

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	clipboard.WriteAll(args)
	fmt.Println(args)
}
