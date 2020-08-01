package backup

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DirHash resumes the files of a path to a final hash.
func DirHash(path string) (finalHash string, err error) {
	hash := md5.New()

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) (err2 error) {
		if err != nil {
			return
		}

		io.WriteString(hash, path)

		fmt.Fprintf(hash, "%v", info.IsDir())
		fmt.Fprintf(hash, "%v", info.ModTime())
		fmt.Fprintf(hash, "%v", info.Mode())
		fmt.Fprintf(hash, "%v", info.Name())
		fmt.Fprintf(hash, "%v", info.Size())
		return
	})
	if err != nil {
		return
	}

	finalHash = fmt.Sprintf("%x", hash.Sum(nil))
	return
}
