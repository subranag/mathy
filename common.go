package mathy

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

var random = rand.New(rand.NewSource(31))
var plusMinus = []int{-1, 1}

func Randomize(point *gg.Point, minx, maxx, miny, maxy float64) {
	point.X = minx + (maxx-minx)*rand.Float64()
	point.Y = miny + (maxy-miny)*random.Float64()
}

func plusOrMinus() int {
	return plusMinus[random.Intn(len(plusMinus))]
}

//NewPoint creates a new point from the given point applying the angle and the scale
//angle is in raidans
func NewPoint(from *gg.Point, angle, scale float64) *gg.Point {
	return &gg.Point{X: from.X + math.Cos(angle)*scale, Y: from.Y + math.Sin(angle)*scale}
}

//NewPointJitter creates a new point from the given point applying the angle and the scale
//angle is in raidans, but with a random jitter
func NewPointJitter(from *gg.Point, angle, scale float64, jitter int) *gg.Point {
	return &gg.Point{
		X: from.X + math.Cos(angle)*scale + float64(plusOrMinus()*rand.Intn(jitter)),
		Y: from.Y + math.Sin(angle)*scale + float64(plusOrMinus()*rand.Intn(jitter))}
}

//RenderInContext creates a canvas and passed the bounds and context to the render function and calls it
func RenderInContext(render func(width, height float64, dc *gg.Context)) {
	// set up everything
	w := 1920
	h := 1080
	width := float64(w)
	height := float64(h)
	dc := gg.NewContext(w, h)
	dc.DrawRectangle(0, 0, width, height)
	dc.SetColor(color.Black)
	dc.Fill()

	// just call the render function
	render(width, height, dc)
	dc.SavePNG("/tmp/out.png")
}

func NewRGBLinearGradient(x0, y0, x1, y1 float64, alpha uint8) *gg.Gradient {
	grad := gg.NewLinearGradient(x0, y0, x1, y1)
	grad.AddColorStop(0, color.NRGBA{0, 255, 0, alpha})
	grad.AddColorStop(0.5, color.NRGBA{255, 0, 0, alpha})
	grad.AddColorStop(1, color.NRGBA{0, 0, 255, alpha})
	return &grad
}
