# uuid-cl

## Description
Simple little program to use for generating different types of unique ids on the command line
* [uuids/guids](https://en.wikipedia.org/wiki/Universally_unique_identifier)
* [cuids](https://usecuid.org/)
* [nanoids](https://github.com/ai/nanoid)
* [ulids](https://github.com/ulid/spec)
* [xids](https://github.com/rs/xid)
* [MongoDB ObjectIDs](https://github.com/mongodb/mongo-go-driver/blob/v1/bson/primitive/objectid.go)


## Why?
I recently discovered [DevToys](https://github.com/veler/DevToys) and was hoping 
the tools would be more command line oriented. So being bored one 
afternoon thought I'd bang out a quick little tool. Yes, there are other
ways to accomplish the basics of generating an uuid. See https://superuser.com/questions/155740/how-can-i-generate-a-uuid-from-the-command-line-in-windows-xp
for ways to do it on Windows. On Linux look into ```uuid``` or ```uuidgen```. On a Mac,  
```uuidgen``` exists.

Another reason for "why," is a simple "because I can." While we should strive to 
follow the "dry" principle, don't be afraid to write little tools that
match _your_ needs.

To further show "why," I expanded the utility to create other types of unique ids in addition to uuids.

## Building

```go build```


## Usage
```uuid <optional args>```

*note: the last generated uuid will be copied to the system clipboard*


#### Optional Arguments

-n=# Number to generate (default = 1) 

-sep=string Separator character for multiples (default = newline)

-l=# length of the id where applicable (modifier to cuids and nanoIds)

-crypt=true|false Generate cryptographic-random ids (default = false, modifier to ulid generation)

-uuid=true|false Generate uuids (default = true if executable is named ```uuid```)

-v=(1|2p|2g|4|6|7) Version to generate (default = 4)

-d=true|false Print dashes in uuids (default = true)

-cuid=true|false Generate cuids (default = false unless executable is named ```cuid```)

-nano=true|false Generate nanoIds (default = false unless executable is named ```nanoid```)

-ulid=true|false Generate nanoIds (default = false unless executable is named ```ulid```)

-xid=true|false Generate xids (default = false unless executable is named ```xid```)

-oid=true|false Generate MongoDB ObjectIDs (default = false unless executable is named ```oid```)

-demo=true|false Generate one of each type (default = false)

(note: on Windows platforms, the executable extension is ignored for comparing executable name)