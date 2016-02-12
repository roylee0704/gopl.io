package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	s = flag.String("s", " ", "seperator")

	n = flag.Bool("n", false, "omit trailing newline")
)

var outVar io.Writer = os.Stdout //modified during testing
func main() {
	flag.Parse()

	// The non-flag arguments are available from flag.Args() as a slice of strings.
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(outVar, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(outVar)
	}
	return nil
}
