
# readme

go中的标准库支持哪些格式:

gif,jpeg, png


## The Image Interface

> 至少从  2015起该接口就没有变了.

```go

type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}

```



base64 是什么情况下:

Instead of writing the image data to a file, 
we could base64 encode it and store it as a string. 
This is useful if you want to generate an image and embed it directly in an HTML document. 
That is beneficial for one-time images that don't need to be stored on the file system and for creating stand-alone HTML documents that don't require a folder full of images to go with it.

> 只要一个字符串，不需要存储的时候用这个比较好.
> 有可能太大了... 比如 opcv-log 有 39k, 出来的字符串比较长。


关于 draw的问题:
```text
// Draw calls DrawMask with a nil mask.
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)

实际使用

```

含义如下:
> https://blog.csdn.net/chenbaoke/article/details/42805627

```text
　  dst  绘图的背景图。
    r 是背景图的绘图区域 是个`巨型`.
    src 是要绘制的图 
    
    sp 是 src 对应的绘图开始点（绘制的大小 r变量定义了）
    mask 是绘图时用的蒙版，控制替换图片的方式。
    mp 是绘图时蒙版开始点（绘制的大小 r变量定义了）
　op Op is a Porter-Duff compositing operator.  参考文章："http://blog.csdn.net/ison81/article/details/5468763"  http://blog.csdn.net/ison81/article/details/5468763 
　Porter-Duff 等式12种规则可以看这篇博客："http://www.blogjava.net/onedaylover/archive/2008/01/16/175675.html" http://www.blogjava.net/onedaylover/archive/2008/01/16/175675.html
```