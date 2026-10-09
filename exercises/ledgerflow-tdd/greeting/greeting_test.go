package greeting

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	table := []struct {
		test string
		name string
		want string
	}{
		{
			test: "When empty returns hello world",
			name: "",
			want: "Hello, world!",
		},
		{
			test: "When have a name returns hello name",
			name: "Rachel",
			want: "Hello, Rachel!",
		},
	}

	for _, v := range table {
		t.Run(v.test, func(t *testing.T) {
			result := Hello(v.name)

			if result != v.want {
				t.Errorf("want %s, have %s", v.want, result)
			}
		})
	}
}
