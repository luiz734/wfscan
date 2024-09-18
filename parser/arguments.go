package parser

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	Path string `short:"p" name:"path" required:"" type:"path" help:"Dir to start scan"`

	Skip    int `short:"m" name:"skip" default:"0" help:"Skip first entries"`
	Entries int `short:"n" name:"entries" default:"20" help:"Number of entries to display"`

	Sort string `short:"s" name:"sortby" default:"size" enum:"size,name"  help:"Define a sorting method"`

	Lesser  string `short:"l" name:"lesser" default:"" help:"Show results lesser than"`
	Greater string `short:"g" name:"greater" default:"" help:"Show results greater than"`

	PrintError bool `short:"e" name:"errors" default:"false" help:"Show error messages"`
}

func foo() {
	_ = kong.Parse
}
