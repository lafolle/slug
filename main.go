package main

import (
	"os"
	"strings"
    "github.com/lafolle/slug/slugify"
)

func main() {
	args := os.Args[1:]
	s := slugify.Slugify(strings.Join(args, " "))
	if len(s) != 0 {
		println(s)
	}
}
