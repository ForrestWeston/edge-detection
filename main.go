package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

type graph struct {
	background color.Color
	elements   []element
	xBound     int
	yBound     int
	NextLabel  int
}

type element struct {
	label int
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

	size := img.Bounds().Max
	background := color.White
	g := new(graph)
	InitalizeGraph(g, img, background)
	fmt.Println("Graph Initalized")

	//Fist pass over image
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			col := img.At(x, y)
			if !sameColor(col, background) {
				neighbors := g.FindNeighbors(x, y)

				if neighbors == nil {
					index := g.GetElementIndexAt(x, y)
					g.elements[index].label = g.NextLabel
					g.NextLabel++
				} else {
					L := neighbors
					index := g.GetElementIndexAt(x, y)
					g.elements[index] = SelectMin(L)
				}

			}
		}
	}

	return
}

func InitalizeGraph(g *graph, i image.Image, bg color.Color) {

	g.background = bg
	g.NextLabel = 0
	for x := 0; x < g.xBound; x++ {
		for y := 0; y < g.yBound; y++ {
			e := new(element)
			if sameColor(i.At(x, y), g.background) {
				e.label = 0
			} else {
				e.label = 1
			}
			g.elements = append(g.elements, *e)
		}
	}
}

func (g *graph) FindNeighbors(x, y int) []element {
	var neighbors []element
	elem := g.GetElementAt(x, y)

	if x == 0 && y == 0 {
		return neighbors
	}
	if y == 0 {
		W := g.GetElementAt(x-1, y)
		if W.label == elem.label {
			neighbors = append(neighbors, W)
		}
		return neighbors
	}
	if x == 0 {
		N := g.GetElementAt(x, y-1)
		NE := g.GetElementAt(x+1, y-1)
		if N.label == elem.label {
			neighbors = append(neighbors, N)
		}
		if NE.label == elem.label {
			neighbors = append(neighbors, NE)
		}
		return neighbors
	}
	NW := g.GetElementAt(x-1, y-1)
	W := g.GetElementAt(x-1, y)
	N := g.GetElementAt(x, y-1)
	NE := g.GetElementAt(x+1, y-1)
	if NW.label == elem.label {
		neighbors = append(neighbors, NW)
	}
	if W.label == elem.label {
		neighbors = append(neighbors, W)
	}
	if N.label == elem.label {
		neighbors = append(neighbors, N)
	}
	if NE.label == elem.label {
		neighbors = append(neighbors, NE)
	}
	return neighbors
}

func (g *graph) GetElementAt(x, y int) element {
	return g.elements[g.xBound*x+y]
}

func (g *graph) GetElementIndexAt(x, y int) int {
	return g.xBound*x + y
}

func SelectMin(elems []element) element {
	e := elems[0]
	for _, oe := range elems[1:] {
		if oe.label < e.label {
			e = oe
		}
	}
	return e
}
