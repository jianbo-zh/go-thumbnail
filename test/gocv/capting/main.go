package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

// 把一个文件的每一帧都给 保存进来，文件名取 数字即可

func main() {

	deviceID := "/Users/apple/Desktop/video/black2.mp4"

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()
	i := 1
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %v\n", deviceID)
			return
		}

		if img.Empty() {
			continue
		}
		outPut := fmt.Sprintf("/Users/apple/workspace_stariverpool/go-image/testdata/output2/%d.jpeg", i)
		gocv.IMWrite(outPut, img)
		i++

	}

}
