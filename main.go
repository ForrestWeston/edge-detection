package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

type walker struct {
	objects []*object
	visited map[image.Point]bool
}

type object struct {
	id    uint32
	c     color.Color
	bound image.Rectangle
}

func (w *walker) Enqueue(o *object) {
	w.objects = append(w.objects, o)
}

func (w *walker) Dequeue() *object {
	o := w.objects[0]
	w.objects = w.objects[1:]
	return o
}

func (w *walker) NumObjects() uint {
	return uint(len(w.objects))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please specify an image")
		return
	}

	var id uint32 = 0

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

	w := new(walker)
	size := img.Bounds().Max
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			//assume that white is the background
			col := img.At(x, y)
			//mark the point as visited
			w.visited[image.Pt(x, y)] = true
			if col != color.White {
				o := new(object)
				o.id = id
				o.c = col
				o.bound.Max.X = x
				o.bound.Max.Y = y
				w.Enqueue(o)
			}
		}
	}
	return
}
