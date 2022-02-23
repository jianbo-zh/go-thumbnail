package main

import (
	"flag"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path"
)

/**
https://www.zhangwenbing.com/blog/golang/Hk-JArrmN



*/

var (
	bg      = flag.String("bg", "bg.png", "背景图片")
	pt      = flag.String("pt", "pt.png", "前景图片") // 水印
	offsetX = flag.Int("offsetX", 0, "x轴偏移值")
	offsetY = flag.Int("offsetY", 0, "y轴偏移值")
	prefix  = flag.String("prefix", "test_", "文件名前缀")
)

func main() {
	flag.Parse()
	mergeImage(*pt)
}

func mergeImage(file string) {
	imgb, _ := os.Open(*bg)    // 背景图
	img, _ := png.Decode(imgb) // 解码
	defer imgb.Close()
	b1 := img.Bounds()

	wmb, _ := os.Open(file)
	watermark, _ := png.Decode(wmb)
	b2 := watermark.Bounds()
	defer wmb.Close()

	offset := image.Pt((b1.Max.X-b2.Max.X)/2+*offsetX, (b1.Max.Y-b2.Max.Y)/2+*offsetY)
	b := img.Bounds()
	m := image.NewRGBA(b) // 同样大小的 新建一个
	draw.Draw(m, b, img, image.Point{}, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.Point{}, draw.Over)

	imgw, _ := os.Create(*prefix + path.Base(file))
	png.Encode(imgw, m)
	defer imgw.Close()
}
