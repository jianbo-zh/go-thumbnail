package image

type FileResult struct {
	Path          string // 文件路径
	Mimetype      string // 文件类型 这个地方就可以判断出该文件的类型是图片/视频/...
	ThumbnailData []byte // 缩略图 的数据 有缩放操作 (图片和视频都有)  缩放的尺寸 (最大不超过120)
	CoverData     []byte // 封面，没有缩放，原尺寸大小 用于视频播放前的预览 (只有视频才有)
}

func Image(filePath string) (*FileResult, error) {

	return nil, nil
}
