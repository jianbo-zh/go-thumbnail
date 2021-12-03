package image

import (
	"fmt"
	"testing"
)

func TestSave2Jpg(t *testing.T) {

	srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"

	f, err := Image(srcFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	thFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo1.jpg"
	coverFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo2.jpg"

	if err = Save2Jpg(f, thFile, coverFile); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("保存成功")

}
