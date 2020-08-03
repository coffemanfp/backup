package backup

import (
	"fmt"
	"path/filepath"
	"time"
)

// Monitor reprensents the monitor data.
type Monitor struct {
	Paths       map[string]string
	Archiver    Archiver
	Destination string
}

// Now starts the monitor
func (m *Monitor) Now() (counter int, err error) {
	var newHash string

	for path, lastHash := range m.Paths {
		newHash, err = DirHash(path)
		if err != nil {
			return
		}

		if newHash != lastHash {
			err = m.act(path)
			if err != nil {
				return
			}
			m.Paths[path] = newHash // update the hash
			counter++
		}
	}

	return
}

func (m *Monitor) act(path string) (err error) {
	dirname := filepath.Base(path)
	filename := fmt.Sprintf("%d.zip", time.Now().UnixNano())

	err = m.Archiver.Archive(path, filepath.Join(m.Destination, dirname, filename))
	return
}
