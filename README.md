# uuid-cl

## Description
Simple little program to use for generating uuids on the 
command line. 

## Why?
I recently discovered https://github.com/veler/DevToys and was hoping 
the tools would be more command line oriented. So being bored one 
afternoon thought I'd bang out a quick little tool. Yes, there are other
ways to accomplish the basics of generating an uuid. See https://superuser.com/questions/155740/how-can-i-generate-a-uuid-from-the-command-line-in-windows-xp
for ways to do it on Windows. On Linux look into ```uuid``` or ```uuidgen```. On a Mac, I 
think ```uuidgen``` exists. I'll try to update this if/when I get access to a Mac again.

Another reason for "why," is a simple "because I can." While we should strive to 
follow the "dry" principle, don't be afraid to write little tools that
match _your_ needs.

## Building

```go build```


## Usage
```uuid <optional args>```


#### Optional Arguments

-n=# Number to generate (default = 1) 

-v=(1|4) Version to generate (default = 4)

-d=true|false Print dashes in uuids (default = true)

-s=string Separator character for multiples (default = newline)
