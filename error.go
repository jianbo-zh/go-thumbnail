package image

import "errors"

//  ErrNotSupportFileCheckMimetype is  returned when the file mimetype  can't be checked
var ErrNotSupportFileCheckMimetype = errors.New("not support file check mimetype") //  未支持的文件类型检查 [文件类型检查的时候失败]

//  ErrNotSupportFile4Img is  returned when can't get a pic from the file, 比如 从 mp3文件里进行截图
var ErrNotSupportFile4Img = errors.New("can't get a pic from this type file")

// ErrNotSupportFileImg is returned when the file is checked but can't  be treated
var ErrNotSupportFileImg = errors.New("not support file type img")

// ErrNotSupportFileVideo is returned when the file is checked but can't  be treated
var ErrNotSupportFileVideo = errors.New("not support file type video")

// ErrNotSupportDirectory is returned when the file is directory
var ErrNotSupportDirectory = errors.New("not support file directory")

// ErrFilePathInvalid is returned when the file path is invalid
var ErrFilePathInvalid = errors.New("the file or directory path is invalid")

// ErrFileNotExist is returned when the file(/directory) path is not exist
var ErrFileNotExist = errors.New("the file or directory is not exist")

// ErrIsNotFile this never happen , everything is file
var ErrIsNotFile = errors.New("err err err !!! this is not  file")

// ErrGoCVInner , for example, gocv inner error is returned when the file is  heif format
var ErrGoCVInner = errors.New("go cv lib is error")

// ErrSave2Jpg , 保存文件出错了
var ErrSave2Jpg = errors.New("save jpg file error")

// /tmp
