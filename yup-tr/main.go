package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tr"
)

const (
	flagDelete     = "delete"
	flagSqueeze    = "squeeze-repeats"
	flagComplement = "complement"
)

func main() {
	app := &cli.App{
		Name:  "tr",
		Usage: "translate or delete characters",
		UsageText: `tr [OPTIONS] SET1 [SET2]

   Translate, squeeze, and/or delete characters from standard input,
   writing to standard output.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagDelete,
				Aliases: []string{"d"},
				Usage:   "delete characters in SET1, do not translate",
			},
			&cli.BoolFlag{
				Name:    flagSqueeze,
				Aliases: []string{"s"},
				Usage:   "replace each sequence of a repeated character with a single occurrence",
			},
			&cli.BoolFlag{
				Name:    flagComplement,
				Aliases: []string{"c"},
				Usage:   "use the complement of SET1",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "tr: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments (SET1 and optionally SET2)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.Bool(flagDelete) {
		params = append(params, Delete)
	}
	if c.Bool(flagSqueeze) {
		params = append(params, Squeeze)
	}
	if c.Bool(flagComplement) {
		params = append(params, Complement)
	}

	// Create and execute the tr command
	cmd := Tr(params...)
	return yup.Run(cmd)
}
