package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Archiver reprensents a archiver handler.
type Archiver interface {
	Archive(src, dest string) error
}

type zipper struct{}

// ZIP is an Archiver that zips and unzips files.
var ZIP Archiver = (*zipper)(nil)

func (z *zipper) Archive(src, dest string) (err error) {
	if err = os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return
	}

	out, err := os.Create(dest)
	if err != nil {
		return
	}
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) (err2 error) {
		if info.IsDir() {
			return // skip
		}
		if err != nil {
			return
		}

		in, err2 := os.Open(path)
		if err != nil {
			return
		}
		defer in.Close()

		f, err2 := w.Create(path)
		if err != nil {
			return
		}

		io.Copy(f, in)
		return
	})

	return
}