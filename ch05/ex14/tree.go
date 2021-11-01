package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileObject struct {
	path string
	name string
	mode int
}

const (
	regular = iota
	directory
	symlink
	block
	character
	socket
	fifo
)

var tree = make(map[string][]FileObject)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(absPath string) []string {
	childDirs := []string{}

	entries, err := os.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		return childDirs
	}

	for _, e := range entries {
		name := e.Name()
		var mode int
		switch e.Type().String()[0] {
		case 'D':
			mode = block
		case 'c':
			mode = character
		case 'd':
			mode = directory
		case 'L':
			mode = symlink
		case 'S':
			mode = socket
		case 'p':
			mode = fifo
		case '-':
			mode = regular
		default:
			log.Fatalf("invalid file mode: %s", e.Type().String())
		}

		obj := FileObject{absPath, name, mode}
		tree[absPath] = append(tree[absPath], obj)
		if obj.mode == directory {
			childDirs = append(childDirs, filepath.Join(absPath, obj.name))
		}
	}
	return childDirs
}

func drawTree(absPath string) {
	indent := "│   "
	branch := "├── "
	indentLast := "    "
	branchLast := "└── "

	printObjName := func(o FileObject) {
		var suffix string
		if o.mode == symlink {
			dst, err := os.Readlink(o.name)
			if err != nil {
				suffix = "[symlink is broken]"
			} else {
				suffix = fmt.Sprintf("-> %s", dst)
			}
		}

		fmt.Printf("%s %s\n", o.name, suffix)
	}

	var draw func(child []FileObject, depth int, prefix string)
	draw = func(child []FileObject, depth int, prefix string) {
		for i, c := range child {
			fmt.Print(prefix)
			if i == len(child)-1 {
				fmt.Print(branchLast)
				printObjName(c)
				if c.mode == directory {
					draw(tree[filepath.Join(c.path, c.name)], depth+1, prefix+indentLast)
				}
			} else {
				fmt.Print(branch)
				printObjName(c)
				if c.mode == directory {
					draw(tree[filepath.Join(c.path, c.name)], depth+1, prefix+indent)
				}
			}
		}
	}

	fmt.Println(filepath.Base(absPath))
	draw(tree[absPath], 1, "")
}

func main() {
	dirname := "."
	if len(os.Args) > 1 {
		dirname = os.Args[1]
	}

	abs, err := filepath.Abs(dirname)
	if err != nil {
		log.Fatal(err)
	}

	breadthFirst(crawl, []string{abs})
	drawTree(abs)
}
