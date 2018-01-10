package cli

import (
	"flag"
	"fmt"
)

var client bool = false
var server bool = false
var clientPortRange []int
var serverPortRange []int

func Help() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	parcePortRange(*wordPtr)
}

func GetTest() string {
	GetFlagsState()

	return ("1111")

}

func GetFlagsState() {
	fmt.Printf("client: %v\n", client)
	fmt.Printf("server: %v\n", server)

}

func parcePortRange(portRange string) {
	for char := range portRange {
		fmt.Printf("%v\n", char)
	}
}
