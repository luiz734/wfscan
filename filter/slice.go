package filter

import (
	"wfscan/parser"
	"wfscan/types"
)

func Slice(files *[]types.FileInfo, cli *parser.CLI) {
    *files = (*files)[cli.Skip:cli.Entries+cli.Skip]
}
