package main

import "testing"

func TestExtract(t *testing.T) {
	cases := []struct {
		Name  string
		Str   string
		Want  string
		IsErr bool
	}{
		{
			Name: "just",
			Str: `
some text before
<CODE>
func main() {
	// some comment
}
</CODE>
some text after
`,
			Want: `
func main() {
	// some comment
}
`,
		},
		{
			Name: "no open tag",
			Str: `
some text before
func main() {
	// some comment
}
</CODE>
some text after
`,
			IsErr: true,
		},
		{
			Name: "no close tag",
			Str: `
some text before
<CODE>
func main() {
	// some comment
}
some text after
`,
			IsErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			code, err := extractTheCode(c.Str)
			if err != nil {
				if c.IsErr {
					return
				}

				t.Fatal(err)
			}

			if code != c.Want {
				t.Logf("is: %s, want: %s", code, c.Want)
				t.Fatal("not equal")
			}
		})
	}
}
