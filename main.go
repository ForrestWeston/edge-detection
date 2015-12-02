package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func sameColor(a, b color.Color) bool {
	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

type walker struct {
	maxSize  image.Point
	minSize  image.Point
	objects  []*object
	explored map[image.Point]bool
	frontier []image.Point
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
	myColor := color.White
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
			if sameColor(myColor, img.At(x+i, y+j)) {
				count++
			}
		}
	}
	if count >= 4 {
		return true
	}
	return false
}

func (w *walker) ExploreObject(o *object, img image.Image, loc image.Point) {
	frontier := []image.Point{}
	start := loc
	myColor := o.c
	Maxsize := img.Bounds().Max
	//Minsize := img.Bounds().Min
	x := start.X
	y := start.Y

	for sameColor(myColor, img.At(x, y)) && x < Maxsize.X {
		w.explored[image.Pt(x, y)] = true
		x++
	}
	x--
	if w.IsVertex(image.Pt(x, y), img) {
		//add the vertex to the objects list
		frontier = append(frontier, image.Pt(x, y))
	}
	x = start.X
	for sameColor(myColor, img.At(x, y)) && y < Maxsize.Y {
		w.explored[image.Pt(x, y)] = true
		y++
	}
	y--
	if w.IsVertex(image.Pt(x, y), img) {
		//add the vertex to the objects list
		frontier = append(frontier, image.Pt(x, y))
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
					o.Enqueue(loc)
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
