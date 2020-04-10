package mashiro

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func ResizeThumbByFile(fileName string) {
	// read file

	//fileByte, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	return
	//}
	//
	//// image decode
	//decodeImage, fileFormatType, err := image.Decode(bytes.NewBuffer(fileByte))
	//if err != nil {
	//	return
	//}
	//
	//// calc image width n height
	//decodeWidth := decodeImage.Bounds().Bounds().Size().X
	//decodeHeight := decodeImage.Bounds().Bounds().Size().Y
	//
	//if decodeWidth - decodeHeight > 0 {
	//	// width pic
	//} else decodeWidth - decodeHeight < 0 {
	//	// height pic
	//} else {
	//	// squre pic
	//}


}

// CutCircle 将一个图按指定像素，以图片重点为原点，画圆切割
func CutCircle() {
	file, err := os.Create("newcircle.png")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	imageFile , err := os.Open("./test_data/3.png")

	if err != nil {
		fmt.Println(err)
	}
	defer imageFile.Close()

	srcImg, _ := png.Decode(imageFile)

	w := srcImg.Bounds().Max.X - srcImg.Bounds().Min.X
	h := srcImg.Bounds().Max.Y - srcImg.Bounds().Min.Y

	d := w
	if w > h {
		d = h
	}

	maskImg := circleMask(d)

	dstImg := image.NewRGBA(image.Rect(0,0,d,d))

	draw.DrawMask(dstImg, srcImg.Bounds().Add(image.Pt(0,0)), srcImg, image.Pt((w-d)/2,(h-d)/2), maskImg,image.Pt(0,0),draw.Src)

	png.Encode(file, dstImg)
}



func circleMask(d int) image.Image{
	img := image.NewRGBA(image.Rect(0,0,d,d))

	for x:= 0;y:=float64(d)/2;{
		img.Set(x,y,color.RGBA{255, 255, 255, 0})
	}else {
		img.Set(x,y,color.RGBA{0, 0, 255, 255})
	}

}
