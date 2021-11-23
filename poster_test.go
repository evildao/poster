package poster

import (
	"bytes"
	"image"
	"image/color"
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/golang/freetype"
)

func ExamplePoster() {
	text := `中国执照的骄傲的司法局噢ID就仨覅哦`
	// 字体装载，只支持ttf，部分不规范的ttf无法解析
	file, _ := ioutil.ReadFile("testdata/FeiHuaSongTi-2.ttf")
	trueTypeFont, _ := freetype.ParseFont(file)
	// 创建画布
	ptx := New(500, 800)
	// 绘制颜色
	ptx.DrawColor(color.RGBA{0xff, 0xff, 0xff, 0xff}, image.Rect(0, 0, 500, 800))
	// 将图片绘制到画布
	file, _ = ioutil.ReadFile("testdata/1.jpg")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(0, 0, 500, 500))
	// 绘制文本
	_ = ptx.DrawText(text, image.Pt(10, 550), 20, color.Gray16{Y: 0}, trueTypeFont)
	// 绘制文本
	_ = ptx.DrawText("￥12557.45", image.Pt(10, 600), 20, color.RGBA{0xff, 0, 0, 0xff}, trueTypeFont)
	// 绘制图片
	file, _ = ioutil.ReadFile("testdata/wechat.jpg")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(-10, -550, 140, -650))
	// 绘制图片
	file, _ = ioutil.ReadFile("testdata/mp.png")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(-200, -550, 20, -650))
	// 输出到文件
	output := ptx.Output()
	_ = ioutil.WriteFile("testdata/out.png", output, fs.ModePerm)
}

func TestNew(t *testing.T) {
	text := `中国执照的骄傲的司法局噢ID就仨覅哦`
	// 字体装载，只支持ttf，部分不规范的ttf无法解析
	file, _ := ioutil.ReadFile("testdata/FeiHuaSongTi-2.ttf")
	trueTypeFont, _ := freetype.ParseFont(file)
	// 创建画布
	ptx := New(500, 800)
	// 绘制颜色
	ptx.DrawColor(color.RGBA{0xff, 0xff, 0xff, 0xff}, image.Rect(0, 0, 500, 800))
	// 将图片绘制到画布
	file, _ = ioutil.ReadFile("testdata/1.jpg")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(0, 0, 500, 500))
	// 绘制文本
	_ = ptx.DrawText(text, image.Pt(10, 550), 20, color.Gray16{Y: 0}, trueTypeFont)
	// 绘制文本
	_ = ptx.DrawText("￥12557.45", image.Pt(10, 600), 20, color.RGBA{0xff, 0, 0, 0xff}, trueTypeFont)
	// 绘制图片
	file, _ = ioutil.ReadFile("testdata/wechat.jpg")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(-10, -550, 140, -650))
	// 绘制图片
	file, _ = ioutil.ReadFile("testdata/mp.png")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(-200, -550, 20, -650))
	// 输出到文件
	output := ptx.Output()
	_ = ioutil.WriteFile("testdata/out.png", output, fs.ModePerm)
}
