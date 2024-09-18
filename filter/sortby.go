package filter

import (
	"slices"
	"strings"
	"wfscan/parser"
	"wfscan/types"
)

type SortFunction func(a, b types.FileInfo) int
var sortby = map[string]SortFunction{
	"size": func(a, b types.FileInfo) int { return int(b.Size - a.Size) },
	"name": func(a, b types.FileInfo) int { return strings.Compare(b.Name, a.Name) },
}

func SortBy(files *[]types.FileInfo, cli *parser.CLI) {
    mapKey := cli.Sort
    sortFn := sortby[mapKey]
	slices.SortFunc(*files, sortFn) 
}
