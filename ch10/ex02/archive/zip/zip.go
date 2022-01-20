package zip

import (
	"archive/zip"

	"github.com/uu64/gpl-book/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat(".zip", &ZipReader{})
}

type ZipReader struct{}

func (zr *ZipReader) Read(filename string) ([]string, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}

	defer r.Close()

	names := []string{}
	for _, f := range r.File {
		names = append(names, f.Name)
	}

	return names, nil
}
