package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
)

func main() {
	var dash bool
	var count int
	var version int
	var sep string
	flag.BoolVar(&dash, "d", true, "Print uuid with dashes")
	flag.IntVar(&count, "n", 1, "Number of uuids generate")
	flag.IntVar(&version, "v", 4, "Version of UUID to generate (1 or 4)")
	flag.StringVar(&sep, "s", "\n", "Separator character for generating multiple uuids")
	flag.Parse()

	for i := 0; i < count; i++ {
		var u uuid.UUID
		if version == 1 {
			u = uuid.Must(uuid.NewUUID())
		} else if version == 4 {
			u = uuid.New()
		} else {
			fmt.Println("Version must be either 1 or 4")
			os.Exit(1)
		}
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
