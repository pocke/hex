package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	flag_reverse  bool
	flag_downcase bool
)

func main() {
	f := flag.NewFlagSet("hex", flag.ExitOnError)
	f.BoolVar(&flag_reverse, "reverse", false, "Hex to Dec")
	f.BoolVar(&flag_downcase, "down", false, "Output by downcase")
	f.Parse(os.Args[1:])

	var fn func(string) (string, error)
	if flag_reverse {
		fn = Dec
	} else {
		fn = Hex
	}

	for _, v := range f.Args() {
		s, err := fn(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(s)
	}
}

func Hex(dec string) (string, error) {
	i, err := strconv.Atoi(dec)
	if err != nil {
		return "", err
	}

	var format string
	if flag_downcase {
		format = "%x"
	} else {
		format = "%X"
	}
	return fmt.Sprintf(format, i), nil
}

func Dec(hex string) (string, error) {
	i, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(i)), nil
}
