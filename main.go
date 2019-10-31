package main

import (
	"fmt"
	"log"
	"os"

	"path/filepath"

	"image"
	"image/color"
	"image/png"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	//nil - Нулевое значение ()
	if err != nil {
		log.Fatal(err)
	}

	files := ""
	fmt.Print("\nВставте имя файла: ")
	fmt.Fscan(os.Stdin, &files)
	file_img := dir + "\\image\\" + files + ".png"
	fmt.Println("file: ", file_img, "\n")

	file, err := os.Open(file_img)
	if err != nil {
		log.Fatal(err)
	}

	// избавляет от необходимости объявлять все переменные для возвращаемых значений
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			R, G, B, _ := img.At(x, y).RGBA()

			//Яркость - это стандартный алгоритм, используемый программным обеспечением для обработки изображений
			//Luma - это похожая форма с гамма-коррекцией, используемая в телевизорах высокой четкости (HDTV):
			//Luma: Y = 0,2126 R + 0,7152 G + 0,0722 B
			Y := (0.2126*float32(R) + 0.7152*float32(G) + 0.0722*float32(B)) * (255.0 / 65535)
			grayPix := color.Gray{uint8(Y)}
			grayImg.Set(x, y, grayPix)
		}
	}

	f, err := os.Create("image\\NewFile.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, grayImg); err != nil {
		log.Fatal(err)
	}
}
