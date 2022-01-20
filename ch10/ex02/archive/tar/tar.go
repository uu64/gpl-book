package tar

import (
	"archive/tar"
	"io"
	"os"

	"github.com/uu64/gpl-book/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat(".tar", &TarReader{})
}

type TarReader struct{}

func (tr *TarReader) Read(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// archive, err := gzip.NewReader(file)
	// if err != nil {
	// 	return nil, err
	// }

	// r := tar.NewReader(archive)
	r := tar.NewReader(file)
	names := []string{}
	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		names = append(names, hdr.Name)
	}
	return names, nil
}
