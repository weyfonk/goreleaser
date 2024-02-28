// Package targz implements the Archive interface providing tar.gz archiving
// and compression.
package targz

import (
	"compress/gzip"
	"io"

	"github.com/weyfonk/goreleaser/pkg/archive/tar"
	"github.com/weyfonk/goreleaser/pkg/config"
)

// Archive as tar.gz.
type Archive struct {
	gw *gzip.Writer
	tw *tar.Archive
}

// New tar.gz archive.
func New(target io.Writer) Archive {
	// the error will be nil since the compression level is valid
	gw, _ := gzip.NewWriterLevel(target, gzip.BestCompression)
	tw := tar.New(gw)
	return Archive{
		gw: gw,
		tw: &tw,
	}
}

func Copying(source io.Reader, target io.Writer) (Archive, error) {
	// the error will be nil since the compression level is valid
	gw, _ := gzip.NewWriterLevel(target, gzip.BestCompression)
	srcgz, err := gzip.NewReader(source)
	if err != nil {
		return Archive{}, err
	}
	tw, err := tar.Copying(srcgz, gw)
	return Archive{
		gw: gw,
		tw: &tw,
	}, err
}

// Close all closeables.
func (a Archive) Close() error {
	if err := a.tw.Close(); err != nil {
		return err
	}
	return a.gw.Close()
}

// Add file to the archive.
func (a Archive) Add(f config.File) error {
	return a.tw.Add(f)
}
