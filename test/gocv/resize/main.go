package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func main() {

	src := gocv.IMRead("/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png", gocv.IMReadColor)
	if src.Empty() {
		fmt.Println("为空数据")
		return
	}
	defer src.Close()

	fmt.Println(src.Rows(), src.Cols())
	fmt.Println(src.Size())

	dst := gocv.NewMat()
	defer dst.Close()

	//gocv.Resize(src, &dst, image.Point{}, 0.5, 0.5, gocv.InterpolationDefault)
	//dst, err := gocv.NewMatFromBytes(200, 356, src.Type(), src.ToBytes())
	//if err != nil {
	//	return
	//}

	croppedMat := src.Region(image.Rect(0, 0, 200, 365))
	dst = croppedMat.Clone()

	if ok := gocv.IMWrite("/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo_out.png", dst); !ok {
		fmt.Println("缩略图 保存失败")
	}
	fmt.Println("保存成功")

}
