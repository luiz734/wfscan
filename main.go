package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"wfscan/filter"
	"wfscan/parser"
	"wfscan/types"

	"github.com/alecthomas/kong"
	// "github.com/dustin/go-humanize"
)

func main() {
	var cli parser.CLI
	ctx := kong.Parse(&cli)

	fmt.Println(ctx.Args)
	var files []types.FileInfo

	filePath := cli.Path

	if err := StatRecursivelly(filePath, &files, 0, cli.PrintError); err != nil {
		errMsg := fmt.Errorf("error stat on filess: %w: ", err)
		fmt.Fprintln(os.Stderr, errMsg)
		os.Exit(1)
	}

	// always sort by size first
	filter.SortBy(&files, &cli)
	filter.KeepSmall(&files, &cli)
	filter.KeepBig(&files, &cli)
	filter.Slice(&files, &cli)

	for _, f := range files {
		fmt.Println(f)
	}

	// fmt.Println(humanize.Bytes(uint64(cli.Greater * 1024)))
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func StatRecursivelly(path string, files *[]types.FileInfo, depth int, debug bool) error {
	// ignore := []string{".git"}
	walkDirFunc := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				if debug {
					log.Printf("skipping dir: %w: ", err.Error())
				}
				return nil
			}
			for _, e := range entries {
				info, err := e.Info()
				if err != nil {
					if debug {
						log.Printf("skipping %w: ", err.Error())
					}
					continue
				}
				*files = append(*files, *types.NewFileInfo(path, info.Name(), info.Size()))
			}
		}
		return nil
	}
	return filepath.WalkDir(path, walkDirFunc)
}
