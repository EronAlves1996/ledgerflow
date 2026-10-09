package greeting

import "fmt"

const (
	empty           = ""
	messageTemplate = "Hello, %s!"
	defaultFiller   = "world"
)

func Hello(name string) string {
	filler := defaultFiller
	if name != empty {
		filler = name
	}

	return fmt.Sprintf(messageTemplate, filler)
}
