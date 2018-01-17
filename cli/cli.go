package cli

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"unicode/utf8"
)

var client bool = false
var server bool = false
var clientPortRange []int
var serverPortRange []int

func Help() {
	wordPtr := flag.String("word", "foo", "a string")

	//numbPtr := flag.Int("numb", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	// fmt.Println("word:", *wordPtr)
	// fmt.Println("numb:", *numbPtr)
	// fmt.Println("fork:", *boolPtr)
	// fmt.Println("svar:", svar)
	// fmt.Println("tail:", flag.Args())
	ParcePortRange(*wordPtr)
}

func GetTest() string {
	GetFlagsState()

	return ""
}

func GetFlagsState() {
	//fmt.Printf("client: %v\n", client)
	//fmt.Printf("server: %v\n", server)

}

func checkPortsSimbol(s string) (string, error) {
	if s == "-" || s == "," {
		return s, nil
	}
	if _, err := strconv.Atoi(s); err != nil {
		return "", err
	}
	return s, nil
}

func chomp(s string) string {
	if len(s) > 0 {
		if s[len(s)-1:len(s)] == "-" || s[len(s)-1:len(s)] == "," {
			s = s[:len(s)-1]
			s = chomp(s)
		}
		if s[0:1] == "-" || s[0:1] == "," {
			s = s[1:]
			s = chomp(s)
		}
		if s[len(s)-1:] == " " {
			s = s[:len(s)-1]
			s = chomp(s)
		}
		if s[0:1] == " " {
			s = s[1:]
			s = chomp(s)
		}
	}
	return s
}

func preparingPortsRange(portRange string) []string {
	var temp string
	var port []string
	count := utf8.RuneCountInString(portRange) - 1
	for i, r := range portRange {
		_, err := strconv.Atoi(string(r))
		if err == nil {
			temp += string(r)
		}
		if (err != nil) || (i == count) {
			if _, err := strconv.Atoi(temp); err == nil || i == count {
				port = append(port, temp)
				//fmt.Printf("%v-\n", test)
				temp = ""
			}
			if string(r) == "-" {
				port = append(port, string(r))
			}

		}
	}
	for i, r := range port {
		//fmt.Println(r)
		if r == "-" {
			port[i], port[i-1] = port[i-1], port[i]
		}
	}
	//fmt.Println(port)
	return port
}
func removeDublicatePorts(ports []int) []int {
	for i := 0; i < len(ports); i++ {
		//fmt.Println(len(ports))
		if i < len(ports)-1 {
			if ports[i] == ports[i+1] {
				//fmt.Printf("%v %v\n", ports[i], ports[i+1])
				temp := ports[:i]
				ports = ports[i+1:]
				ports = append(temp, ports...)

				//fmt.Printf("%v \n", ports)
				i--
			}
		}

	}
	//fmt.Println(ports)
	return ports
}

func createPortRange(ports []string) ([]int, error) {
	var from, to int
	var err error
	var portInt []int
	for i := 0; i < len(ports); {
		if ports[i] == "-" {
			if i < len(ports) {
				if from, err = strconv.Atoi(string(ports[i+1])); err != nil {
					return portInt, fmt.Errorf("Wrong port num: %v", from)
				}
				if to, err = strconv.Atoi(string(ports[i+2])); err != nil {
					return portInt, fmt.Errorf("Wrong port num: %v", to)
				}
				temp := ports[:i]
				ports = ports[i+3:]
				ports = append(temp, ports...)
				i--
			}
			if from < to && from > 0 && to <= 65535 {
				portInt = append(portInt, makePortRange(from, to)...)
			} else {
				return portInt, fmt.Errorf("Wrong port range num: %v-%v", from, to)
			}
		} else {
			if from, err = strconv.Atoi(string(ports[i])); err != nil {
				return portInt, fmt.Errorf("Wrong port num: %v", from)
			}
			if from > 65535 {
				return portInt, fmt.Errorf("Wrong port num: %v", from)
			}
			portInt = append(portInt, from)
		}
		from = 0
		to = 0
		i++
	}
	sort.Ints(portInt)
	//fmt.Printf("%v \n", portInt)
	portInt = removeDublicatePorts(portInt)
	//fmt.Printf("%v \n", portInt)
	return portInt, nil
}
func ParcePortRange(portRange string) ([]int, error) {
	var ports []string
	var portInt []int
	portRange = chomp(portRange)
	//fmt.Printf("%q \n", string(portRange))
	var err error
	for _, r := range portRange {
		//fmt.Printf("%q \n", string(i))
		_, err := checkPortsSimbol(string(r))
		//fmt.Printf("%v \n", err)
		if err != nil {
			return portInt, fmt.Errorf("Wrong value of parameter: %q", portRange)
		}
	}
	//fmt.Printf("%v \n", ports)
	ports = preparingPortsRange(portRange)
	//fmt.Printf("%v \n", ports)
	portInt, err = createPortRange(ports)
	//fmt.Printf("%v \n", portInt)
	//fmt.Printf("%v \n", err)
	return portInt, err
}

func makePortRange(from, to int) []int {
	var port []int
	if from < to {
		for i := from; i <= to; i++ {
			port = append(port, i)
		}
	}
	//fmt.Printf("%v \n", port)
	return port
}
