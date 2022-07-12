package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/jaevor/go-nanoid"
	"github.com/lucsky/cuid"
	"os"
	"strings"
)

func main() {
	var dash bool
	var version int

	var slug bool
	var crypt bool

	var genCuid bool
	var genNanoid bool

	var count int
	var sep string

	var nanoLength int

	flag.BoolVar(&slug, "slug", false, "Generate a cuid slug instead of uuid")
	flag.BoolVar(&crypt, "crypt", false, "Generate cryptographic random cuid instead of uuid")
	flag.BoolVar(&dash, "d", true, "Print uuid with dashes")
	flag.IntVar(&version, "v", 4, "Version of UUID to generate (1 or 4)")
	flag.BoolVar(&genCuid, "cuid", false, "Generate cuid instead of uuid")
	flag.BoolVar(&genNanoid, "nano", false, "Generate nanoid instead of uuid")
	flag.IntVar(&count, "n", 1, "Number to generate")
	flag.StringVar(&sep, "sep", "\n", "Separator character to use when generating multiples")
	flag.IntVar(&nanoLength, "l", 21, "Length of a nanoid to generate")
	flag.Parse()

	for i := 0; i < count; i++ {
		var u string
		if genCuid || slug || crypt {
			u = createCuid(slug, crypt)
		} else if genNanoid {
			u = createNanoid(nanoLength)
		} else {
			u = createUUID(dash, version)
		}

		lastChar := sep
		if i == count-1 {
			lastChar = ""
		}

		fmt.Printf("%s%s", u, lastChar)
	}
}

func createUUID(dash bool, version int) string {
	var u uuid.UUID
	if version == 1 {
		u = uuid.Must(uuid.NewUUID())
	} else if version == 4 {
		u = uuid.New()
	} else {
		fmt.Println("Version must be either 1 or 4")
		os.Exit(2)
	}
	if dash {
		return fmt.Sprintf("%v", u)
	} else {
		return fmt.Sprintf("%s", strings.Replace(u.String(), "-", "", -1))
	}
}

func createCuid(slug bool, crypt bool) string {
	var c string
	if crypt {
		var err error
		c, err = cuid.NewCrypto(rand.Reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
	} else {
		if slug {
			c = cuid.Slug()
		} else {
			c = cuid.New()
		}
	}
	return c
}

func createNanoid(length int) string {
	generator, err := nanoid.Standard(length)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	return generator()
}
