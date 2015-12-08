package main

import (
	"fmt"
	"image"
	"image/color"
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

	w := &walker{
		explored: make(map[image.Point]bool),
	}
	size := img.Bounds().Max
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			//assume that white is the background
			col := img.At(x, y)
			loc := image.Pt(x, y)
			w.explored[loc] = true
			if !sameColor(col, color.White) {
				if w.IsVertex(loc, img) {
					fmt.Println("Found Vertex", x, y)
					o := new(object)
					//add the vertex to the objects list
					o.EnqueueVertex(loc)
					o.c = col
					//increase x look for vert, increase y look for vert
					w.ExploreObject(o, img, loc)
					fmt.Println(o.verticies)
				}
			}
		}
	}
	fmt.Println(w.NumObjects())
	fmt.Println(w.objects)
	return
}
