package main

import (
	"image"
	"image/color"

	_ "image/jpeg"
	_ "image/png"
)

type object struct {
	verticies []image.Point
	c         color.Color
	bound     image.Rectangle
}

func (o *object) EnqueueVertex(p image.Point) {
	o.verticies = append(o.verticies, p)
}

func (o *object) DequeueVertex() image.Point {
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
