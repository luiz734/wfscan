package parser

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	Path string `short:"p" name:"path" required:"" type:"path" help:"Dir to start scan"`

	Skip    int `short:"m" name:"skip" default:"0" help:"Skip first entries"`
	Entries int `short:"n" name:"entries" default:"20" help:"Number of entries to display"`

	Sort string `short:"s" name:"sortby" default:"size" enum:"size,name"  help:"Define a sorting method"`

	Smaller  string `short:"l" name:"smaller" default:"" help:"Show results smaller than"`
	Larger string `short:"g" name:"larger" default:"" help:"Show results larger than"`

	PrintError bool `short:"e" name:"errors" default:"false" help:"Show error messages"`
	// Format     Format `short:"o" name:"format default:"sf" help:"Output format"`
	Format string `short:"o" name:"format" default:"sf" help:"Output format:\n\ts: file size\n\tf: file full path\n\td: file dir\n\tn: file name\n"`
}

func foo() {
	_ = kong.Parse
}
