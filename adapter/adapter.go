package main

import (
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}
type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{Lines: []Line{
		Line{0, 0, width, 0},
		Line{0, 0, width, height},
		Line{width, 0, width, height},
		Line{0, height, width, height},
	}}
}

type Point struct {
	X, Y int
}
type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}

	}
	maxX += 1
	maxY += 1
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}
	for _, point := range points {
		data[point.Y][point.X] = '*'
	}
	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}
	return b.String()
}

type vectorToRasterAdapter struct {
	points []Point
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func (v vectorToRasterAdapter) addLine(line Line) {

}
func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return adapter
}

func main() {

}
