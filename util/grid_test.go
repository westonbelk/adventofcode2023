package util

import (
	"image"
	"reflect"
	"testing"
	"unicode"
)

const gridIn = `..c..
...#.
.2..1`

func TestGrid(t *testing.T) {
	type args struct {
		in     []byte
		filter GridFilter
	}
	tests := []struct {
		name string
		args args
		want map[image.Point]rune
	}{
		{
			name: "default read grid",
			args: args{in: []byte(gridIn), filter: func(r rune) bool { return true }},
			want: map[image.Point]rune{
				{0, 0}: '.',
				{1, 0}: '.',
				{2, 0}: 'c',
				{3, 0}: '.',
				{4, 0}: '.',

				{0, 1}: '.',
				{1, 1}: '.',
				{2, 1}: '.',
				{3, 1}: '#',
				{4, 1}: '.',

				{0, 2}: '.',
				{1, 2}: '2',
				{2, 2}: '.',
				{3, 2}: '.',
				{4, 2}: '1',
			},
		},
		{
			name: "default read grid with filter",
			args: args{in: []byte(gridIn), filter: func(r rune) bool { return unicode.IsDigit(r) }},
			want: map[image.Point]rune{
				{1, 2}: '2',
				{4, 2}: '1',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Grid(tt.args.in, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdjacentToPoint(t *testing.T) {
	type args struct {
		p image.Point
	}
	tests := []struct {
		name string
		args args
		want []image.Point
	}{
		{
			name: "{0,0}",
			args: args{p: image.Point{0, 0}},
			want: []image.Point{
				{-1, -1}, {-1, 0}, {-1, 1},
				{0, -1}, {0, 0}, {0, 1},
				{1, -1}, {1, 0}, {1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjacentToPoint(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdjacentToPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdjacentToRect(t *testing.T) {
	type args struct {
		r image.Rectangle
	}
	tests := []struct {
		name string
		args args
		want []image.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjacentToRect(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdjacentToRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectPoints(t *testing.T) {
	type args struct {
		rect image.Rectangle
	}
	tests := []struct {
		name string
		args args
		want []image.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RectPoints(tt.args.rect); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
