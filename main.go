package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/jaevor/go-nanoid"
	"github.com/lucsky/cuid"
	"github.com/oklog/ulid/v2"
	"os"
	"strings"
	"time"
)

func main() {
	var dash bool
	var version int

	var slug bool
	var crypt bool

	var genUuid bool
	var genCuid bool
	var genNanoid bool
	var genUlid bool

	var count int
	var sep string

	var length int

	flag.BoolVar(&genUuid, "uuid", false, "Generate uuid")
	flag.BoolVar(&dash, "d", true, "Print uuid with dashes")
	flag.IntVar(&version, "v", 4, "Version of UUID to generate (1 or 4)")
	flag.BoolVar(&genCuid, "cuid", false, "Generate cuid")
	flag.BoolVar(&slug, "slug", false, "Generate a slug (modifier to cuid)")
	flag.BoolVar(&crypt, "crypt", false, "Generate cryptographic strong id (modifier to cuid and ulid)")
	flag.BoolVar(&genNanoid, "nano", false, "Generate nanoid")
	flag.BoolVar(&genUlid, "ulid", false, "Generate ulid")
	flag.IntVar(&count, "n", 1, "Number to generate")
	flag.StringVar(&sep, "sep", "\n", "Separator character to use when generating multiples")
	flag.IntVar(&length, "l", 0, "Length of a unique id to generate")
	flag.Parse()

	if !(genUuid || genCuid || genNanoid || genUlid) {
		appName := strings.ToLower(os.Args[0])
		if strings.HasPrefix(appName, "uuid") {
			genUuid = true
		} else if strings.HasPrefix(appName, "cuid") {
			genCuid = true
		} else if strings.HasPrefix(appName, "nanoid") {
			genNanoid = true
		} else if strings.HasPrefix(appName, "ulid") {
			genUlid = true
		}
	}

	for i := 0; i < count; i++ {
		var u string
		if genCuid {
			u = createCuid(slug, crypt)
		} else if genNanoid {
			u = createNanoid(length)
		} else if genUuid {
			u = createUUID(dash, version)
		} else if genUlid {
			u = createUlid(crypt)
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
	l := length
	if l < 1 {
		length = 21
	}
	generator, err := nanoid.Standard(l)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	return generator()
}

func createUlid(crypt bool) string {
	if crypt {
		entropy := rand.Reader
		ms := ulid.Timestamp(time.Now())
		id, err := ulid.New(ms, entropy)
		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		return id.String()
	}
	// default to fast/less secure
	return ulid.Make().String()
}
