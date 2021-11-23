// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poster_test

import (
	"bytes"
	"image"
	"image/color"
	"io/fs"
	"io/ioutil"
	"t/poster"

	"github.com/golang/freetype"
)

func ExamplePoster() {
	text := `中国执照的骄傲的司法局噢ID就仨覅哦`

	file, _ := ioutil.ReadFile("testdata/FeiHuaSongTi-2.ttf")
	trueTypeFont, _ := freetype.ParseFont(file)

	ptx := poster.New(500, 800)

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
	_ = ioutil.WriteFile("testdata/out.png", output, fs.ModePerm)
}
