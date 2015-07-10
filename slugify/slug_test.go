package slugify

import "testing"

func TestSlugify(t *testing.T) {
    cases := []struct {
            in, want string
    }{
            {"How are you doing", "how-are-you-doing"},
            {"this", "this"},
            {"  what is  in the name   ", "what-is-in-the-name"},
            {"", ""},
    }
    for _, c := range cases {
        got := Slugify(c.in)
        if got != c.want {
            t.Errorf("Slugify(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
