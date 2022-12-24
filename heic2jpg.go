package image

import (
	"image/jpeg"
	"io"
	"log"
	"os"

	"github.com/adrium/goheif"
)

func init() {
	// 容器需要设置安全编码否则会崩溃
	goheif.SafeEncoding = true
}

// Skip Writer for exif writing
type writerSkipper struct {
	w           io.Writer
	bytesToSkip int
}

func (w *writerSkipper) Write(data []byte) (int, error) {
	if w.bytesToSkip <= 0 {
		return w.w.Write(data)
	}

	if dataLen := len(data); dataLen < w.bytesToSkip {
		w.bytesToSkip -= dataLen
		return dataLen, nil
	}

	if n, err := w.w.Write(data[w.bytesToSkip:]); err == nil {
		n += w.bytesToSkip
		w.bytesToSkip = 0
		return n, nil
	} else {
		return n, err
	}
}

func newWriterExif(w io.Writer, exif []byte) (io.Writer, error) {
	writer := &writerSkipper{w, 2}
	soi := []byte{0xff, 0xd8}
	if _, err := w.Write(soi); err != nil {
		return nil, err
	}

	if exif != nil {
		app1Marker := 0xe1
		markerlen := 2 + len(exif)
		marker := []byte{0xff, uint8(app1Marker), uint8(markerlen >> 8), uint8(markerlen & 0xff)}
		if _, err := w.Write(marker); err != nil {
			return nil, err
		}

		if _, err := w.Write(exif); err != nil {
			return nil, err
		}
	}

	return writer, nil
}

func HeicConvert2jpg(fileIn string, fileOut string) error {

	fi, err := os.Open(fileIn)
	if err != nil {
		return err
	}
	defer fi.Close()

	exif, err := goheif.ExtractExif(fi)
	if err != nil {
		log.Printf("Warning: no EXIF from %s: %v\n", fileIn, err)
	}

	img, err := goheif.Decode(fi)
	if err != nil {
		log.Printf("Failed to parse %s: %v\n", fileIn, err)
		return err
	}

	fo, err := os.OpenFile(fileOut, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Failed to create output file %s: %v\n", fileOut, err)
		return err
	}
	defer fo.Close()

	w, err := newWriterExif(fo, exif)
	if err != nil {
		log.Printf("new writer exif error: %v\n", err)
		return err
	}

	err = jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Printf("Failed to encode %s: %v\n", fileOut, err)
		return err
	}

	log.Printf("Convert %s to %s successfully\n", fileIn, fileOut)
	return nil
}
