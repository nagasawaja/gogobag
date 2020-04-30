package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("newcircle12.png")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	dstImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	writeOnImage(dstImg, 24, fmt.Sprintf("￥%s", "123"), color.RGBA{177, 122, 0, 127}, 30, 30)

	png.Encode(file, dstImg)
}

func writeOnImage(target *image.RGBA, fontSize float64, fontString string, fontColor color.Color, x, y int) {
	c := freetype.NewContext()

	c.SetClip(target.Bounds())
	c.SetDst(target)
	c.SetHinting(font.HintingVertical)

	// 设置文字颜色、字体、字大小
	c.SetSrc(image.NewUniform(fontColor))
	c.SetFontSize(fontSize)
	fontFam, err := getFontFamily()
	if err != nil {
		log.Println("get font family error")
	}
	c.SetFont(fontFam)
	pt := freetype.Pt(x, y)

	_, err = c.DrawString(fontString, pt)
	if err != nil {
		log.Printf("draw error: %v \n", err)
	}

}

func getFontFamily() (*truetype.Font, error) {

	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := ioutil.ReadFile("msyhbd.ttf")
	if err != nil {
		log.Println("read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}
