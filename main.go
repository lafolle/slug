package main

import (
	"fmt"
	"os"
	"strings"
)

func Slugify(s string) string {
	words := strings.Split(s, " ")
	var slug, w string
	for i := range words {
		w = strings.ToLower(strings.TrimSpace(words[i]))
		if len(w) == 0 {
			continue
		}
		if i == 0 {
			slug = w
		} else {
			slug = fmt.Sprintf("%s-%s", slug, w)
		}
	}
	return slug
}

func main() {
	args := os.Args[1:]
	slug := Slugify(strings.Join(args, " "))
	if len(slug) != 0 {
		println(slug)
	}
}
