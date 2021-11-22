package main

import (
	"image"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
)
var text = `中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符中文字符`
func main() {
	ptx := New(500, 800)

	ptx.DrawColor(color.RGBA{0xff, 0xff, 0xff, 0xff}, image.Rect(0, 0, 500, 800))

	// file, err := ioutil.ReadFile("testdata/2.png")
	// log.Println(err)
	// err = ptx.DrawImage(bytes.NewReader(file), image.Rect(0, 0, 500, 500))
	// log.Println(err)

	file, _ := ioutil.ReadFile("testdata/font.ttf")
	trueTypeFont, err := freetype.ParseFont(file)
	log.Println(err)
	err = ptx.DrawText(text, image.Pt(-100, 100), 20, color.RGBA{0xff, 0xff, 0xff, 0xff}, trueTypeFont)
	log.Println(err)

	output := ptx.Output()
	ioutil.WriteFile("out.png", output, 0644)
}
