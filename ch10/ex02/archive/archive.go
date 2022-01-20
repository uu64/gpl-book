package archive

import (
	"fmt"
	"path/filepath"
)

type ArchiveReader interface {
	Read(filename string) ([]string, error)
}

var readers map[string]ArchiveReader = make(map[string]ArchiveReader)

func RegisterFormat(ext string, ar ArchiveReader) {
	readers[ext] = ar
}

func Read(filename string) ([]string, error) {
	ext := filepath.Ext(filename)

	if _, ok := readers[ext]; !ok {
		return nil, fmt.Errorf("reader not registered: %s", ext)
	}

	return readers[ext].Read(filename)
}
