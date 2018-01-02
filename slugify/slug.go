package slugify

import (
	"fmt"
	"strings"
)

func Slugify(s string) string {
	s = strings.TrimSpace(s)
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
