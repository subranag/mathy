package mathy

import (
	"math"

	"github.com/fogleman/gg"
)

func Sierpinski() {

	RenderInContext(func(width, height float64, dc *gg.Context) {

		jitter := 10
		layers := 3
		sizef := float64(0.25)
		alpha := uint8(30)

		for i := 0; i < layers; i++ {
			factor := height * sizef
			drawSierpinskiLayer(dc, factor, 0, 0, width, height, jitter, alpha)
			jitter -= 3
			alpha += 30
			sizef += 0.05
		}

		center := &gg.Point{X: width * 0.5, Y: height * 0.65}
		factor := height * 0.6
		grad := NewRGBLinearGradient(0, 0, width, height, 255)
		dc.SetStrokeStyle(*grad)
		dc.SetLineWidth(2)
		sierpinskiRecurse(dc, center, 7, factor, 3)
		dc.Stroke()
	})
}

func drawSierpinskiLayer(dc *gg.Context, factor, x0, y0, x1, y1 float64, jitter int, alpha uint8) {
	grad2 := NewRGBLinearGradient(x0, y0, x1, y1, alpha)
	dc.SetStrokeStyle(*grad2)
	dc.SetLineWidth(1)

	point := &gg.Point{}
	for i := 0; i < 30; i++ {
		Randomize(point, x0, x1, y0, y1)
		sierpinskiRecurse(dc, point, 5, factor, jitter)
	}
	dc.Stroke()
}

func sierpinskiRecurse(dc *gg.Context, origin *gg.Point, maxDepth int, scale float64, jitter int) {
	if maxDepth == 0 {
		drawTriangle(dc, origin, scale, jitter)
		return
	}

	newScale := scale / 2
	angle := gg.Radians(-90)
	top := NewPointJitter(origin, angle, newScale, jitter)
	sierpinskiRecurse(dc, top, maxDepth-1, newScale, jitter)

	angle += gg.Radians(120)
	right := NewPointJitter(origin, angle, newScale, jitter)
	sierpinskiRecurse(dc, right, maxDepth-1, newScale, jitter)

	angle += gg.Radians(120)
	left := NewPointJitter(origin, angle, newScale, jitter)
	sierpinskiRecurse(dc, left, maxDepth-1, newScale, jitter)
}

func drawTriangle(dc *gg.Context, origin *gg.Point, scale float64, jitter int) {

	angle := -1 * math.Pi / 2
	top := NewPointJitter(origin, angle, scale, jitter)
	dc.MoveTo(top.X, top.Y)

	angle += math.Pi * 2 / 3
	right := NewPointJitter(origin, angle, scale, jitter)
	dc.LineTo(right.X, right.Y)

	angle += math.Pi * 2 / 3
	left := NewPointJitter(origin, angle, scale, jitter)
	dc.LineTo(left.X, left.Y)

	dc.LineTo(top.X, top.Y)
}
