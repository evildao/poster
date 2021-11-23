package poster

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"

	"github.com/anthonynsimon/bild/transform"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Poster interface {
	DrawColor(color color.Color, rt image.Rectangle)
	DrawImage(r io.Reader, rt image.Rectangle) error
	DrawText(text string, pt image.Point, size float64, color color.Color, tf *truetype.Font) error
	Output() []byte
}

type posterImpl struct {
	img  *image.RGBA
	font *truetype.Font
}

func New(width, height int) Poster {
	posterInstance := &posterImpl{
		img: image.NewRGBA(image.Rect(0, 0, width, height)),
	}
	return posterInstance
}

// DrawColor 绘制颜色
func (p *posterImpl) DrawColor(color color.Color, rt image.Rectangle) {
	for x := rt.Min.X; x <= rt.Max.X; x++ {
		for y := rt.Min.Y; y <= rt.Max.Y; y++ {
			p.img.Set(x, y, color)
		}
	}
}

// DrawImage 绘制图片
func (p *posterImpl) DrawImage(r io.Reader, rt image.Rectangle) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	if !rt.Empty() && !img.Bounds().Eq(rt) {
		img = transform.Resize(img, rt.Dx(), rt.Dy(), transform.Linear)
	}
	draw.Draw(p.img, p.img.Bounds(), img, rt.Min, draw.Over)
	return nil
}

// DrawText 绘制文字
func (p *posterImpl) DrawText(text string, pt image.Point, size float64, color color.Color, tf *truetype.Font) error {
	fc := freetype.NewContext()
	// 设置屏幕每英寸的分辨率
	fc.SetDPI(72)
	// 设置用于绘制文本的字体
	fc.SetFont(tf)
	// 以磅为单位设置字体大小
	fc.SetFontSize(size)
	// 设置剪裁矩形以进行绘制
	fc.SetClip(p.img.Bounds())
	// 设置目标图像
	fc.SetDst(p.img)
	// 设置绘制操作的源图像，通常为 image.Uniform
	fc.SetSrc(image.NewUniform(color))
	// fc.SetSrc(image.Black)
	fc.SetHinting(font.HintingFull)

	fpt := freetype.Pt(pt.X, pt.Y)
	_, err := fc.DrawString(text, fpt)
	if err != nil {
		return err
	}
	return nil
}

// Output 输出
func (p *posterImpl) Output() []byte {
	wr := bytes.NewBuffer([]byte{})
	if err := png.Encode(wr, p.img); err != nil {
		log.Fatal(err)
	}
	return wr.Bytes()
}
