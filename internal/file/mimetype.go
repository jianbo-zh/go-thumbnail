package file

import (
	"github.com/gabriel-vasile/mimetype"
)

func DetectFile(filePath string) (*mimetype.MIME, error) {

	return mimetype.DetectFile(filePath)
}
