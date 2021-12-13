package image

import (
	"fmt"
	"image"
	"io"
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

		dst := gocv.NewMat()
		defer dst.Close() //  FIXME: 清空指针和数据 ,一旦清空，返回值就没有了, 如果不清空，内存会存在泄漏

		// 缩略图
		gocv.Resize(src, &dst, image.Point{}, 0.25, 0.25, gocv.InterpolationDefault) // 按原比例缩放， 缩放后的长 ,宽为 原来的25%
		//gocv.Resize(src, &dst, image.Pt(120, 68), 0, 0, gocv.InterpolationDefault) //两种缩放方式

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

		//
		for i := 0; i < 200; i++ {

			ok := webcam.Read(&img)
			if ok {

				if !img.Empty() {

					// 跳过第50帧，也就是 防止黑色
					if i <= 50 {
						continue
					}

					r.CoverData = img

					dst := gocv.NewMat()

					// 缩略图
					gocv.Resize(img, &dst, image.Point{}, 0.25, 0.25, gocv.InterpolationDefault)
					//gocv.Resize(img, &dst, image.Pt(120, 68), 0, 0, gocv.InterpolationDefault) //两种缩放方式
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

// DetectFileMime 根据文件输入路径进行检查 文件类型
func DetectFileMimeByPath(fileInPath string) (string, error) {
	if fileInPath == "" || strings.TrimSpace(fileInPath) == "" {
		return "", ErrFilePathInvalid
	}

	if !file.Exists(fileInPath) {
		return "", ErrFileNotExist
	}

	if file.IsDir(fileInPath) {
		return "", ErrNotSupportDirectory
	}

	if !file.IsFile(fileInPath) {
		return "", ErrIsNotFile // 应该不会发生
	}

	m, err := file.DetectFile(fileInPath)
	if err != nil {
		return "", ErrNotSupportFileCheckMimetype // 不支持的文件类型检查 ,只有检查到了才可以
	}
	return m.String(), nil

}

// DetectFileMimeByByte 可直接传递 文件切片的一部分内容 ,例如 文件开始的 256个字节
func DetectFileMimeByByte(src []byte) (string, error) {
	m, err := file.Detect(src)
	if err != nil {
		return "", ErrNotSupportFileCheckMimetype // 不支持的文件类型检查 ,只有检查到了才可以
	}
	return m.String(), nil
}

// DetectReader 从 reader 里读取 ,返回文件类型
func DetectReader(r io.Reader) (string, error) {
	m, err := file.DetectReader(r)
	if err != nil {
		return "", err
	}
	return m.String(), nil
}
