package image


import "errors"

//  ErrNotSupportFileMimetype is  returned when the file mimetype  can't be checked
var ErrNotSupportFileMimetype = errors.New("not support file mimetype") //  未支持的文件类型


// ErrNotSupportFileImg is returned when the file is checked but can't  treated by opencv
var ErrNotSupportFileImg = errors.New("not support file type img")

// ErrNotSupportFileVideo is returned when the file is checked but can't  treated by FFMpeg
var ErrNotSupportFileVideo = errors.New("not support file type video")



