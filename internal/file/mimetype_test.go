package file

import (
	"fmt"
	"testing"

	"github.com/gabriel-vasile/mimetype"
)

func TestDetectFile(t *testing.T) {

	//filePath := "/Users/apple/workspace_stariverpool/go-image/testdata/other/mp3/1.mp3"
	//filePath := "/Users/apple/workspace_stariverpool/go-image/testdata/other/mp4/1.mp4"
	//filePath := "/Users/apple/workspace_stariverpool/go-image/testdata/other/mp4/2.mp4"

	// 测试图片类型
	//filePath := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	//filePath := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv/apple.jpg"
	filePath := "/Users/apple/workspace/go_concurrency/doc/concurrency_in_go/Concurrency-in-Go.pdf"
	m, err := DetectFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(m.String(), m.Extension(), m.Parent().String())

	fmt.Println(mimetype.EqualsAny(m.String(), "image/png", "image/jpeg"))
	//mimetype.Lookup()
}
