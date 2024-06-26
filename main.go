package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jaevor/go-nanoid"
	"github.com/nrednav/cuid2"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	var dash bool
	var version int

	var crypt bool

	var genUuid bool
	var genCuid bool
	var genNanoid bool
	var genUlid bool
	var genXid bool
	var genObjectId bool
	var demo bool

	var count int
	var sep string

	var length int

	flag.BoolVar(&demo, "demo", false, "Generate one of each")
	flag.BoolVar(&genUuid, "uuid", false, "Generate uuid")
	flag.BoolVar(&dash, "d", true, "Print uuid with dashes")
	flag.IntVar(&version, "v", 4, "Version of UUID to generate (1, 4, 6, or 7)")
	flag.BoolVar(&genCuid, "cuid", false, "Generate cuid")
	flag.BoolVar(&crypt, "crypt", false, "Generate cryptographic strong id (modifier to ulid)")
	flag.BoolVar(&genNanoid, "nano", false, "Generate nanoid")
	flag.BoolVar(&genUlid, "ulid", false, "Generate ulid")
	flag.BoolVar(&genXid, "xid", false, "Generate xid")
	flag.BoolVar(&genObjectId, "oid", false, "Generate MongoDB ObjectID")
	flag.IntVar(&count, "n", 1, "Number to generate")
	flag.StringVar(&sep, "sep", "\n", "Separator character to use when generating multiples")
	flag.IntVar(&length, "l", 0, "Length of a unique id to generate (modifier to cuid and nanoid")
	flag.Parse()

	if demo {
		fmt.Printf("uuid: %s\n", createUUID(dash, version))
		fmt.Printf("cuid: %s\n", createCuid(length))
		fmt.Printf("nanoid: %s\n", createNanoid(length))
		fmt.Printf("ulid: %s\n", createUlid(crypt))
		fmt.Printf("xid: %s\n", createXid())
		fmt.Printf("mongodb ObjectID: %s\n", createObjectID())
		os.Exit(0)
	}

	types := 0
	if genUuid {
		types++
	}
	if genCuid {
		types++
	}
	if genNanoid {
		types++
	}
	if genUlid {
		types++
	}
	if genXid {
		types++
	}
	if genObjectId {
		types++
	}

	if types > 1 {
		fmt.Fprintln(os.Stderr, "Can only create one type of identifier at a time")
		os.Exit(1)
	}

	if types == 0 {
		appName := strings.ToLower(os.Args[0])
		if strings.Contains(appName, "uuid") {
			genUuid = true
		} else if strings.Contains(appName, "cuid") {
			genCuid = true
		} else if strings.Contains(appName, "nanoid") {
			genNanoid = true
		} else if strings.Contains(appName, "ulid") {
			genUlid = true
		} else if strings.Contains(appName, "xid") {
			genXid = true
		} else if strings.Contains(appName, "oid") {
			genObjectId = true
		}
	}

	for i := 0; i < count; i++ {
		var u string
		if genCuid {
			u = createCuid(length)
		} else if genNanoid {
			u = createNanoid(length)
		} else if genUuid {
			u = createUUID(dash, version)
		} else if genUlid {
			u = createUlid(crypt)
		} else if genXid {
			u = createXid()
		} else if genObjectId {
			u = createObjectID()
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
	} else if version == 6 {
		u = uuid.Must(uuid.NewV6())
	} else if version == 7 {
		u = uuid.Must(uuid.NewV7())
	} else {
		fmt.Fprintln(os.Stderr, "Version must be either 1, 4, 6, or 7")
		os.Exit(2)
	}
	if dash {
		return fmt.Sprintf("%v", u)
	} else {
		return strings.ReplaceAll(u.String(), "-", "")
	}
}

func createCuid(length int) string {
	if length < 2 || length > 32 {
		return cuid2.Generate()
	}
	generate, err := cuid2.Init(
		cuid2.WithLength(length),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	return generate()
}

func createNanoid(length int) string {
	l := length
	if l < 1 {
		l = 21
	}
	generator, err := nanoid.Standard(l)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
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
			fmt.Fprintln(os.Stderr, err)
			os.Exit(5)
		}
		return id.String()
	}
	// default to fast/less secure
	return ulid.Make().String()
}

func createXid() string {
	return xid.New().String()
}

func createObjectID() string {
	return primitive.NewObjectID().Hex()
}
