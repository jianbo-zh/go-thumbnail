package image

import (
	"fmt"
	"testing"
)

func TestSave2Jpg(t *testing.T) {

	srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"

	// 只截图，不保存
	f, err := Image(srcFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	thFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo1.jpg"
	coverFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo2.jpg"

	// 保存到 指定目录
	if err = Save2Jpg(f, thFile, coverFile); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("保存成功")

}

func TestImageAndSave(t *testing.T) {
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"

	//srcFile := "/Users/apple/Desktop/video/bafybeic5x7c6bv56t7kg25y57ohmm7ffnfdmleq423ln3xvh4xtll5tpjm.mp4"
	srcFile := "/Users/apple/Desktop/video/black2.mp4"
	// 只截图，不保存
	f, err := ImageAndSave(srcFile, "/Users/apple/workspace_stariverpool/go-image/testdata/output/")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.ThumbnailImgPath, f.CoverImgPath)

}
