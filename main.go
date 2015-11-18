package main

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please specify an image")
		return
	}

	fi, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	img, typ, err := image.Decode(fi)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("image was a ", typ)

	size := img.Bounds().Max
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := img.At(x, y)
			fmt.Printf("color at %d %d:", x, y)
			fmt.Println(color.RGBA())
		}
	}

}
