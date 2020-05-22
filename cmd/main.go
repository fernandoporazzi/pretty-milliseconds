package main

import (
	"fmt"

	prettyms "github.com/fernandoporazzi/pretty-milliseconds"
)

func main() {
	fmt.Println(prettyms.Humanize(649993232323, prettyms.Options{
		Verbose: true,
	}))
}
