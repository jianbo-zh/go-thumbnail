package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

/**

export CGO_CXXFLAGS="--std=c++11"
export CGO_CPPFLAGS="-I/usr/local/opencv452/include"
export CGO_LDFLAGS="-L/usr/local/opencv452/lib -lopencv_stitching -lopencv_superres -lopencv_videostab -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_photo -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_optflow -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_dnn -lopencv_plot -lopencv_shape -lopencv_video -lopencv_ml -lopencv_ximgproc -lopencv_calib3d -lopencv_features2d -lopencv_highgui -lopencv_videoio -lopencv_flann -lopencv_xobjdetect -lopencv_imgcodecs -lopencv_objdetect -lopencv_xphoto -lopencv_imgproc -lopencv_core"

*/

// 拿到 []byte 能否保存进某个图片
func main() {

	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())

	filename := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	img := gocv.IMRead(filename, gocv.IMReadAnyColor)

	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return
	}

	//fmt.Println(img.ToBytes())

	// 保存图片
	//gocv.IMWrite("/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo2.png", img)

	//gocv.NewMatFromBytes()

	//img.Type()
	//img.Size()
}
