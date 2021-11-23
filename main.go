package poster

import (
	"bytes"
	"image"
	"image/color"
	"io/fs"
	"io/ioutil"

	"github.com/golang/freetype"
)

var text = `中国执照的骄傲的司法局噢ID就仨覅哦`

func main() {
	file, _ := ioutil.ReadFile("testdata/FeiHuaSongTi-2.ttf")
	trueTypeFont, _ := freetype.ParseFont(file)

	ptx := New(500, 800)

	ptx.DrawColor(color.RGBA{0xff, 0xff, 0xff, 0xff}, image.Rect(0, 0, 500, 800))

	file, _ = ioutil.ReadFile("testdata/1.jpg")

	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(0, 0, 500, 500))

	_ = ptx.DrawText(text, image.Pt(10, 550), 20, color.Gray16{Y: 0}, trueTypeFont)

	_ = ptx.DrawText("￥12557.45", image.Pt(10, 600), 20, color.RGBA{0xff, 0, 0, 0xff}, trueTypeFont)

	file, _ = ioutil.ReadFile("testdata/wechat.jpg")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(-10, -550, 140, -650))

	file, _ = ioutil.ReadFile("testdata/mp.png")
	_ = ptx.DrawImage(bytes.NewReader(file), image.Rect(-200, -550, 20, -650))

	output := ptx.Output()
	_ = ioutil.WriteFile("out.png", output, fs.ModePerm)
}
