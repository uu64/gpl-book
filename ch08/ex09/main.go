package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type fileSize struct {
	root string
	size int64
}

type diskUsage struct {
	nfiles int64
	nbytes int64
}

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan fileSize)
	usage := make(map[string]*diskUsage)

	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
		usage[root] = &diskUsage{0, 0}
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case fs, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			usage[fs.root].nfiles++
			usage[fs.root].nbytes += fs.size
		case <-tick:
			printDiskUsage(usage)
		}
	}

	printDiskUsage(usage) // final totals
}

func printDiskUsage(usage map[string]*diskUsage) {
	for k, v := range usage {
		fmt.Printf("%s: %d files\t%.1f GB\n", k, v.nfiles, float64(v.nbytes)/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(root, dir string, n *sync.WaitGroup, fileSizes chan<- fileSize) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, fileSizes)
		} else {
			fileSizes <- fileSize{
				root: root,
				size: entry.Size(),
			}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
