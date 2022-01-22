package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main() {
	var dash bool
	var count int
	var sep string
	flag.BoolVar(&dash, "d", true, "Print uuid with dashes")
	flag.IntVar(&count, "n", 1, "Number of uuids generate")
	flag.StringVar(&sep, "s", "\n", "Separator character for generating multiple uuids")
	flag.Parse()

	for i := 0; i < count; i++ {
		u := uuid.New()
		lastChar := sep
		if i == count-1 {
			lastChar = ""
		}
		if dash {
			fmt.Printf("%v%s", u, lastChar)
		} else {
			fmt.Printf("%s%s", strings.Replace(u.String(), "-", "", -1), lastChar)
		}
	}
}
