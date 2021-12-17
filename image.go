package image

import (
	"fmt"
	"image"
	"path"
	"strings"

	"gocv.io/x/gocv"

	"github.com/penglonghua/go-image/internal/file"
)

type FileResult struct {
	SrcPath  string // 文件路径
	Mimetype string // 文件类型 这个地方就可以判断出该文件的类型是图片/视频/...
	//ThumbnailData []byte // 缩略图 的数据 有缩放操作 (图片和视频都有)  缩放的尺寸 (最大不超过120)
	//CoverData     []byte // 封面，没有缩放，原尺寸大小 用于视频播放前的预览 (只有视频才有)
	ThumbnailData    gocv.Mat // 缩略图 的数据 有缩放操作 (图片和视频都有)  缩放的尺寸 (最大不超过120)
	CoverData        gocv.Mat // 封面，没有缩放，原尺寸大小 用于视频播放前的预览 (只有视频才有)
	OutputDir        string   // 输出的文件目录  传入的参数
	ThumbnailImgPath string   // 缩略图 的数据 存放路径
	CoverImgPath     string   // 封面 的数据 存放路径

}

func Image(filePath string) (*FileResult, error) {

	if filePath == "" || strings.TrimSpace(filePath) == "" {
		return nil, ErrFilePathInvalid
	}

	if !file.Exists(filePath) {
		return nil, ErrFileNotExist
	}

	if file.IsDir(filePath) {
		return nil, ErrNotSupportDirectory
	}

	if !file.IsFile(filePath) {
		return nil, ErrIsNotFile // 应该不会发生
	}

	// 确定了一个已经存在的文件
	m, err := file.DetectFile(filePath)
	if err != nil {
		return nil, ErrNotSupportFileCheckMimetype // 不支持的文件类型检查 ,只有检查到了才可以
	}

	// 下面有3种
	// 1 是 图片, 2 是视频， 3 是其他的
	mimeType := m.String()

	if strings.HasPrefix(mimeType, "image") {
		// opencv截图 缩略图
		src := gocv.IMRead(filePath, gocv.IMReadColor)
		if src.Empty() {
			return nil, ErrGoCVInner
		}

		dst := gocv.NewMat()
		//defer dst.Close() // 清空指针和数据 ,一旦清空，返回值就没有了, 如果清空，内存会存在泄漏

		// 缩略图
		gocv.Resize(src, &dst, image.Point{}, 0.25, 0.25, gocv.InterpolationDefault)
		//gocv.Resize(src, &dst, image.Pt(120, 68), 0, 0, gocv.InterpolationDefault) //两种缩放方式

		r := &FileResult{
			SrcPath:       filePath,
			Mimetype:      mimeType,
			ThumbnailData: dst,
			//CoverData:     dst,
		}

		//fmt.Println("是否有数据: ", len(dst.ToBytes()))

		return r, nil

	} else if strings.HasPrefix(mimeType, "video") {
		// 视频 封面 +  缩略图
		webcam, err := gocv.VideoCaptureFile(filePath)
		if err != nil {
			fmt.Printf("Error opening video capture device: %v\n", filePath)
			return nil, ErrGoCVInner
		}
		defer webcam.Close()

		img := gocv.NewMat()
		//defer img.Close()

		r := &FileResult{
			SrcPath:  filePath,
			Mimetype: mimeType,
			//ThumbnailData: &dst,
			//CoverData:     nil,
		}

		for {
			ok := webcam.Read(&img)
			if ok {
				if !img.Empty() {

					r.CoverData = img

					dst := gocv.NewMat()

					// 缩略图
					gocv.Resize(img, &dst, image.Point{}, 0.25, 0.25, gocv.InterpolationDefault)
					r.ThumbnailData = dst // 缩略图
					r.CoverData = img     // 原图

					//dst.Close()

					break
				}
			}
		}

		return r, nil

	} else {
		// 其他的不支持
		return nil, ErrNotSupportFile4Img // 不能从该文件中获取到 图片, 比如 从mp3文件里，是截不了图的
	}

	return nil, nil // 这个地方应该不会执行到
}

// Save2Jpg 保存进文件
func Save2Jpg(f *FileResult, thumbnailSaveFile string, coverSaveFile string) error {

	if f == nil {
		return ErrSave2Jpg
	}

	if strings.HasPrefix(f.Mimetype, "image") {
		if ok := gocv.IMWrite(thumbnailSaveFile, f.ThumbnailData); !ok {
			return ErrSave2Jpg
		}
	} else if strings.HasPrefix(f.Mimetype, "video") {
		if ok := gocv.IMWrite(thumbnailSaveFile, f.ThumbnailData); !ok {
			return ErrSave2Jpg
		}
		if ok := gocv.IMWrite(coverSaveFile, f.CoverData); !ok {
			return ErrSave2Jpg
		}
	} else {
		return ErrNotSupportFile4Img
	}

	// 使用完后之后，需要关闭 ，否则会有内存泄漏
	defer Close(f)

	return nil

}

func Close(f *FileResult) {

	if f != nil {

		f.ThumbnailData.Close()
		f.CoverData.Close()

	}

}

// ImageAndSave 该方法结合了 Image 和 Save2Jpg ,调用者不需要关心内部使用细节
// 根据原始文件路径，在 输出目录下生成对应的 缩略图和封面图片
func ImageAndSave(fileInPath string, outputDir string) (*FileResult, error) {

	//1 输入参数校验

	if fileInPath == "" || strings.TrimSpace(fileInPath) == "" {
		return nil, ErrFilePathInvalid
	}

	if !file.Exists(fileInPath) {
		return nil, ErrFileNotExist
	}

	if file.IsDir(fileInPath) {
		return nil, ErrNotSupportDirectory
	}

	if !file.IsFile(fileInPath) {
		return nil, ErrIsNotFile // 应该不会发生
	}

	// 文件夹判断
	if outputDir == "" || strings.TrimSpace(outputDir) == "" {
		outputDir = file.GetSystemTmp()
	} else {
		// 如果不为空的的情况下

		if !file.Exists(outputDir) {
			return nil, ErrFileNotExist
		}

		if !file.IsDir(outputDir) {
			return nil, ErrFilePathInvalid
		}

	}

	m, err := file.DetectFile(fileInPath)
	if err != nil {
		return nil, ErrNotSupportFileCheckMimetype // 不支持的文件类型检查 ,只有检查到了才可以
	}

	// 下面有3种
	// 1 是 图片, 2 是视频， 3 是其他的 分别做不同的处理
	mimeType := m.String()

	if strings.HasPrefix(mimeType, "image") {
		// opencv截图 缩略图
		src := gocv.IMRead(fileInPath, gocv.IMReadColor)
		if src.Empty() {
			fmt.Printf("Error read the file: %v\n", fileInPath)
			return nil, ErrGoCVInner // FIXME 有可能因为没有响应的解码器 而出错
		}

		defer src.Close() // FIXME 注意回收

		dst := gocv.NewMat()
		defer dst.Close() //  FIXME: 清空指针和数据 ,一旦清空，返回值就没有了, 如果不清空，内存会存在泄漏

		// 缩略图
		//gocv.Resize(src, &dst, image.Point{}, 0.05, 0.05, gocv.InterpolationDefault) // 图片是 0.05
		//gocv.Resize(src, &dst, image.Pt(120, 68), 0, 0, gocv.InterpolationDefault) //两种缩放方式
		Resize(src, dst)

		destThumbnailPath := path.Join(outputDir, "xhh.jpg") // 特定的文件名

		if ok := gocv.IMWrite(destThumbnailPath, dst); !ok {
			return nil, ErrSave2Jpg
		}

		r := &FileResult{
			SrcPath:       fileInPath,
			Mimetype:      mimeType,
			ThumbnailData: dst,
			//CoverData:     dst,
			ThumbnailImgPath: destThumbnailPath,
			CoverImgPath:     "", // 图片没有封面，只有缩略图
		}

		return r, nil

	} else if strings.HasPrefix(mimeType, "video") {
		// 视频 封面 +  缩略图
		webcam, err := gocv.VideoCaptureFile(fileInPath)
		if err != nil {
			fmt.Printf("Error opening video capture device: %v\n", fileInPath)
			return nil, ErrGoCVInner
		}
		defer webcam.Close()

		img := gocv.NewMat()
		defer img.Close()

		r := &FileResult{
			SrcPath:  fileInPath,
			Mimetype: mimeType,
			//ThumbnailData: &dst,
			//CoverData:     nil,
		}

		for i := 0; i < 200; i++ {

			ok := webcam.Read(&img)
			if ok {

				if !img.Empty() {

					// 跳过第20帧，也就是 防止黑色 (具体应该跳过多少帧，应该有一个图像质量判断的方法.)
					if i <= 20 {
						continue
					}

					r.CoverData = img

					dst := gocv.NewMat()

					// 缩略图
					//gocv.Resize(img, &dst, image.Point{}, 0.25, 0.25, gocv.InterpolationDefault) // 视频暂时确定是 0.25
					//gocv.Resize(img, &dst, image.Pt(120, 68), 0, 0, gocv.InterpolationDefault) //两种缩放方式

					Resize(img, dst)

					r.ThumbnailData = dst // 缩略图
					r.CoverData = img     // 原图

					destThumbnailPath := path.Join(outputDir, "xhh.jpg") // 特定的文件名

					if ok := gocv.IMWrite(destThumbnailPath, dst); !ok {
						continue
					}
					r.ThumbnailImgPath = destThumbnailPath

					destCoverPath := path.Join(outputDir, "xcc.jpg") // 特定的文件名
					if ok := gocv.IMWrite(destCoverPath, img); !ok {
						continue
					}
					r.CoverImgPath = destCoverPath

					dst.Close()

					break
				}
			}

		}

		return r, nil

	} else {
		// 其他的不支持
		return nil, ErrNotSupportFile4Img // 不能从该文件中获取到 图片, 比如 从mp3文件里，是截不了图的
	}

	return nil, nil // 这个地方应该不会执行到

}

// 图片缩放要求:
// 对于一般图片，比例没有非常失调的 情况下， 最短一边保持 200,另外一边保持常宽比例不变
// 对于特殊比例失调的图片，处理下 最短一遍也是 200, 最常一边 从中切图 (保持 长的部分为 9:16 的比例， 那么就是  356)
//  比例失调定义为 最大边/最短边 >= 4倍以上
func Resize(src gocv.Mat, dst gocv.Mat) {

	const (
		Min         = 200 // 最短一边是 200
		DefaultMax  = 356 // 为  9:16 = 200:356
		DivideTimes = 4.0 // 最长一边/最短一边的 比例， 超过此比例，定义为 失调 4 倍定义为
	)

	//Width Height
	srcWidth := src.Cols()
	srcHeight := src.Rows()

	srcMax, srcMin := srcWidth, srcHeight
	var horizontal bool = true // 默认是 横的  // （vertical  竖）
	if srcWidth < srcHeight {
		srcMax, srcMin = srcHeight, srcWidth
		horizontal = false // 竖的 照片
	}

	// 获取 最大值，和最小值
	// 调试信息
	//fmt.Printf("宽: %d , 高 : %d\n", srcWidth, srcHeight)
	//fmt.Printf("最大值: %d , 最小值 : %d\n", srcMax, srcMin)
	//if horizontal {
	//	fmt.Println("横图")
	//} else {
	//	fmt.Println("竖图")
	//}

	var maxDividemin float64 = float64(srcMax) / float64(srcMin)
	//fmt.Printf("最长边/最短边 比例 %v\n", maxDividemin)
	if maxDividemin >= DivideTimes {
		//fmt.Println("比例失调")
		// 需要保持 9:16 的比例 , 也就是  200: 365 (最短一边是  200, 最长一一遍是  365)

		if srcMax <= DefaultMax { //  最长一遍小于 365， 无法处理, 直接返回 原始图片
			dst = src
		} else {

			x, y := Min, DefaultMax // 竖的照片
			if horizontal {
				x, y = DefaultMax, Min
			}

			// --------------
			// FIXME: 潜在问题，该方法返回后 数据丢失问题 后面可以在继续测试一下 为什么会丢的问题 ,单个方法不会丢，方法组合在一起使用会丢失
			//fmt.Println("需要截取的 矩型 ", x, y)
			//croppedMat := src.Region(image.Rect(0, 0, x, y))
			//dst = croppedMat.Clone() // 如果原始的 src 丢失这个地方也失效了
			////dst = croppedMat // 如果原始的 src 丢失这个地方也失效了
			////dst, _ = gocv.NewMatFromBytes(croppedMat.Rows(), croppedMat.Cols(), croppedMat.Type(), croppedMat.ToBytes())
			//fmt.Println("处理后的图片", len(dst.ToBytes()))
			// ----------------

			gocv.Resize(src, &dst, image.Pt(x, y), 0, 0, gocv.InterpolationCubic)

		}

	} else {
		// 比例不失调
		x, y := srcWidth, srcHeight

		if horizontal { // 横的照片

			y = Min //  必须先确定
			x = int(float64(y) * maxDividemin)

		} else { // 竖的照片
			x = Min //  必须先确定
			y = int(float64(x) * maxDividemin)
		}
		gocv.Resize(src, &dst, image.Pt(x, y), 0, 0, gocv.InterpolationCubic)
	}

}
