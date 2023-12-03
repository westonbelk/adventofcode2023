package util

import (
	"image"
	"strings"
)

type GridFilter func(r rune) bool

func Grid(in []byte, filter GridFilter) map[image.Point]rune {
	g := make(map[image.Point]rune)
	for y, s := range strings.Fields(string(in)) {
		for x, r := range s {
			if filter(r) {
				g[image.Point{x, y}] = r
			}
		}
	}
	return g
}

func AdjacentToPoint(p image.Point) []image.Point {
	return RectPoints(image.Rectangle{
		Min: image.Point{p.X - 1, p.Y - 1},
		Max: image.Point{p.X + 1, p.Y + 1},
	}.Canon())
}

func AdjacentToRect(r image.Rectangle) []image.Point {
	set := make(map[image.Point]struct{})

	points := RectPoints(r)
	for _, p := range points {
		for _, a := range AdjacentToPoint(p) {
			set[a] = struct{}{}
		}
	}
	res := make([]image.Point, 0, len(set))
	for item := range set {
		res = append(res, item)
	}
	return res
}

func RectPoints(rect image.Rectangle) []image.Point {
	points := make([]image.Point, 0)
	rect = rect.Canon()
	m := rect.Min
	for x := m.X; x <= m.X+rect.Dx(); x++ {
		for y := m.Y; y <= m.Y+rect.Dy(); y++ {
			points = append(points, image.Point{x, y})
		}
	}
	return points
}
