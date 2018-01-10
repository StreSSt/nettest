package main

import (
	"fmt"

	"github.com/StreSSt/nettest/cli"
)

func main() {
	// Parce options
	cli.Help()
	fmt.Printf(cli.GetTest() + "\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
}
