package main

import (
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func main() {
	file, err := os.Create("newcircle1.png")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	imageFile, err := os.Open("4.png")

	if err != nil {
		fmt.Println(err)
	}
	defer imageFile.Close()

	srcImg, _ := png.Decode(imageFile)

	dstImg, _ := Circle(srcImg)

	png.Encode(file, dstImg)
}

func Circle(src image.Image) (*image.RGBA, error) {

	weight := src.Bounds().Max.X
	height := src.Bounds().Max.Y

	sideLen := 0
	if weight >= height {
		sideLen = height
	} else {
		sideLen = weight
	}

	maskImg := buildMask(sideLen)

	dstImg := image.NewRGBA(image.Rect(0, 0, sideLen, sideLen))

	draw.DrawMask(dstImg, src.Bounds().Bounds(), src, image.Pt((weight-sideLen)/2, (height-sideLen)/2), maskImg, image.Pt(0, 0), draw.Src)

	return dstImg, nil
}

func buildMask(d int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, d, d))
	// 计算出原点位置
	originX, originY := math.Ceil(float64(d)/2), math.Ceil(float64(d)/2)
	// 半径
	radius := math.Pow(math.Ceil(float64(d/2)-2), 2)
	//
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			// 直角三角形的两边边长
			aSide := float64(x) - originX
			bSIde := float64(y) - originY
			if aSide*aSide+bSIde*bSIde > radius {
				// 圆外，包括边缘
				if aSide*aSide+bSIde*bSIde >= radius+float64(d) {
					// 边缘
					img.Set(x, y, color.RGBA{255, 255, 255, 0})
				} else if aSide*aSide+bSIde*bSIde >= radius+float64(d)/2 {
					img.Set(x, y, color.RGBA{0, 0, 0, 50})
				} else {
					img.Set(x, y, color.RGBA{255, 255, 255, 170})
				}
			} else {
				// 圆内
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	return img
}
