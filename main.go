package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	nFlag := flag.Int("n", 0, "n")

	flag.Parse()

	var n int
	var err error

	if *nFlag != 0 {
		n = *nFlag
	} else if len(flag.Args()) > 0 {
		n, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			fmt.Println("Invalid number provided:", flag.Arg(0))
			return
		}
	}

	sum := 0

	for i := 0; i <= n; i++ {
		sum += i

	}

	println(sum)

	// fmt.Println(n)
	// utils.MagicChange(&n)
	// fmt.Println(n)

	// magicNumber := utils.MagicNumber{Number: n}
	// fmt.Println("Number Before Multiply: ", magicNumber.Number)
	// magicNumber.Multiply(n)
	// fmt.Println("Number After Multiply: ", magicNumber.Number)

	// fmt.Println(utils.MagicTria(n))

}
