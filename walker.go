package main

import (
	"image"
	"image/color"

	_ "image/jpeg"
	_ "image/png"
)

type walker struct {
	maxSize  image.Point
	minSize  image.Point
	objects  []*object
	explored map[image.Point]bool
	frontier []image.Point
}

func (w *walker) EnqueueObject(o *object) {
	w.objects = append(w.objects, o)
}

func (w *walker) DequeueObject() *object {
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
