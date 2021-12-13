package file

import (
	"errors"
	"io"

	"github.com/gabriel-vasile/mimetype"
)

func DetectFile(filePath string) (*mimetype.MIME, error) {

	return mimetype.DetectFile(filePath)
}

func Detect(src []byte) (*mimetype.MIME, error) {
	if src == nil || len(src) == 0 {
		return nil, errors.New("src is null")
	}
	return mimetype.Detect(src), nil

}

func DetectReader(r io.Reader) (*mimetype.MIME, error) {
	return mimetype.DetectReader(r)
}
