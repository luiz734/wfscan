package printer

import (
	"fmt"
	"path/filepath"
	"strings"
	"wfscan/parser"
	"wfscan/types"

	"github.com/dustin/go-humanize"
)

type FormatFunction func(a types.FileInfo) string

var formats = map[rune]FormatFunction{
	's': func(f types.FileInfo) string { return humanize.Bytes(f.Size) },
	'f': func(f types.FileInfo) string { return filepath.Join(f.Dir, f.Name) },
	'd': func(f types.FileInfo) string { return f.Dir },
	'n': func(f types.FileInfo) string { return f.Name },
}

func Format(f types.FileInfo, cli *parser.CLI) (string, error) {
	var sb strings.Builder
	for _, s := range cli.Format {
		formatFn, ok := formats[s]
		if !ok {
			return "", fmt.Errorf("unknown format option `%s`", string(s))
		}
		sb.WriteString(formatFn(f))
		sb.WriteRune('\t')
	}
	return sb.String(), nil
}
