# go-image

go-image

1. 通过图片文件或者视频文件，从该文件中生产jpg格式的图或者缩略图, 如果该文件为 非图片非视频，那么返回错误,提示不支持该文件类型.
2. 前期支持 输入 常规的 图片，视频格式 ,对 HEIF(.heic)图片格式 暂时不考虑.




一般情况下 只要 opencv和FFMpeg 能够处理，这个地方就能处理。 
## 支持的图片


## 暂时不支持的图片格式 

* HEIF(.heic)  （后期版本）


## 支持的视频

## 不支持的视频



## 环境问题

这个地方推荐使用

CPU版本

```text
docker pull gocv/opencv:4.5.4

```



gpu版本 cuda11, cuda10 需要根据 响应的 显卡环境来.
```text

docker pull gocv/opencv:4.5.4-gpu-cuda-11
docker pull gocv/opencv:4.5.4-gpu-cuda-10

```

how to test in docker.
```shell
docker run -v /Users/apple/workspace_stariverpool/go-image:/Users/apple/workspace_stariverpool/go-image -it   gocv/opencv:4.5.4 bash

cd /Users/apple/workspace_stariverpool/go-image
go mod tidy

go test -v -run=TestImageAndSave  .

```


这个地方还得检查一下.
```shell
git tag -a v0.0.8
git commit 
git push
git push --tags

```

