package main

import (
	"fmt"

	image "github.com/jinabo-zh/go-thumbnail"
)

func main() {

	srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	//srcFile := "/Users/apple/workspace_stariverpool/go-image/testdata/other/mp4/1.mp4"
	//srcFile := "testdata/other/heif/bafybeievjz3aiygjtz343alfiypforxfjeikmt3q4cnek3dtz2jo6s2epy.heic"

	f, err := image.Image(srcFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.SrcPath, f.Mimetype)

	fmt.Println("是否有数据222 : ", len(f.ThumbnailData.ToBytes()))

	// 用完之后再关闭
	//defer f.ThumbnailData.Close() // 应该关闭这个，否则会有内存泄漏

	fmt.Println("是否有数据333 : ", f.ThumbnailData) // 没数据了

	//数据
	//fmt.Println(f.ThumbnailData.ToBytes()) // 这个地方为什么为空的
	//fmt.Println((*f.ThumbnailData).ToBytes())

	thFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo1.jpg"
	coverFile := "/Users/apple/workspace_stariverpool/go-image/testdata/output/opencv-logo2.jpg"
	//
	if err = image.Save2Jpg(f, thFile, coverFile); err != nil {
		fmt.Println(err)
		return
	}

}
