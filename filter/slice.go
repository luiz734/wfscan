package filter

import (
	"wfscan/parser"
	"wfscan/types"
)

func Slice(files *[]types.FileInfo, cli *parser.CLI) {
    startIndex := max(0, cli.Skip)
    endIndex := min(cli.Entries+cli.Skip, len(*files))
    *files = (*files)[startIndex:endIndex]
}
