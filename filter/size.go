package filter

import (
	"fmt"
	"wfscan/parser"
	"wfscan/types"

	"github.com/dustin/go-humanize"
)

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

// type SortFunction func(a, b types.FileInfo) int
// var sortby = map[string]SortFunction{
// 	"size": func(a, b types.FileInfo) int { return int(b.Size - a.Size) },
// 	"name": func(a, b types.FileInfo) int { return strings.Compare(b.Name, a.Name) },
// 	"size-r": func(a, b types.FileInfo) int { return int(a.Size - b.Size) },
// 	"name-r": func(a, b types.FileInfo) int { return strings.Compare(a.Name, b.Name) },
// }

func KeepBig(files *[]types.FileInfo, cli *parser.CLI) error {
	if cli.Larger == "" {
		return nil
	}
	minBytesAllowed, err := humanize.ParseBytes(cli.Larger)
	if err != nil {
		return fmt.Errorf("invalid size: %w : ", err)
	}
	var out []types.FileInfo
	for _, f := range *files {
		if f.Size >= minBytesAllowed {
			out = append(out, f)
		}
	}
	// no no no
	// files = &out

	// yes yes yes
	*files = out
    return nil
}

func KeepSmall(files *[]types.FileInfo, cli *parser.CLI) error {
	if cli.Smaller == "" {
		return nil
	}
	maxBytesAllowed, err := humanize.ParseBytes(cli.Smaller)
	if err != nil {
		return fmt.Errorf("invalid size: %w : ", err)
	}
	var out []types.FileInfo
	for _, f := range *files {
		if f.Size <= maxBytesAllowed {
			out = append(out, f)
		}
	}
	// no no no
	// files = &out

	// yes yes yes
	*files = out
    return nil
}
