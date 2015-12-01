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
	verticies []image.Point
	c         color.Color
	bound     image.Rectangle
}

func (o *object) Enqueue(p image.Point) {
	o.verticies = append(o.verticies, p)
}

func (o *object) Dequeue() image.Point {
	v := o.verticies[0]
	o.verticies = o.verticies[1:]
	return v
}

func (o *object) NumObjects() uint {
	return uint(len(o.verticies))
}

func (o *object) Peek() image.Point {
	return o.verticies[0]
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

func (w *walker) IsVertex(cell image.Point, img image.Image) bool {
	// a point is a vertex iff it has 4 'different' neightbors, this includes diagonal
	count := 0
	x := cell.X
	y := cell.Y
	myColor := img.At(x, y)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			Maxsize := img.Bounds().Max
			Minsize := img.Bounds().Min
			if x+i > Maxsize.X || x+i < Minsize.X {
				count++
			}
			if y+i > Maxsize.Y || y+i < Minsize.Y {
				count++
			}
			if myColor != img.At(x+i, y+j) {
				count++
			}
		}
	}
	if count >= 4 {
		return true
	}
	return false
}

func (w *walker) ExploreObject(o *object, img image.Image) {
	start := o.Peek()
	myColor := o.c
	x := start.X
	y := start.Y

	for myColor == img.At(x, y) {
		x++
	}
}

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

	w := new(walker)
	size := img.Bounds().Max
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			//assume that white is the background
			col := img.At(x, y)
			loc := image.Pt(x, y)
			w.visited[loc] = true
			if col != color.White {
				if w.IsVertex(loc, img) {
					o := new(object)
					//add the vertex to the objects list
					o.Enqueue(loc)
					o.c = col
					w.ExploreObject(o, img)

				}
			}
		}
	}
	return
}
