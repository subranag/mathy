package mathy

import (
	"math"

	"github.com/fogleman/gg"
)

func Koch() {

	RenderInContext(func(width, height float64, dc *gg.Context) {

		grad2 := NewRGBLinearGradient(0, 0, width, height, 255)
		dc.SetStrokeStyle(*grad2)
		dc.SetLineWidth(1.5)
		p0 := &gg.Point{X: 0.3 * width, Y: 0.3 * height}
		p1 := &gg.Point{X: 0.7 * width, Y: 0.3 * height}
		dist := Distance(p0, p1)
		p2 := NewPoint(p0, gg.Radians(60), dist)

		jitter := 1
		kochRecursive(dc, p0, p1, 4, jitter)
		kochRecursive(dc, p1, p2, 4, jitter)
		kochRecursive(dc, p2, p0, 4, jitter)
	})
}

func kochRecursive(dc *gg.Context, p0, p1 *gg.Point, maxDepth, jitter int) {
	dist := Distance(p0, p1)
	split := dist / 3
	dx := p1.X - p0.X
	dy := p1.Y - p0.Y
	angle := math.Atan2(dy, dx)

	pA := NewPointJitter(p0, angle, split, jitter)
	pB := NewPointJitter(pA, angle-gg.Radians(60), split, jitter)
	pC := NewPointJitter(p0, angle, 2*split, jitter)

	if maxDepth == 0 {
		dc.DrawLine(p0.X, p0.Y, pA.X, pA.Y)
		dc.DrawLine(pA.X, pA.Y, pB.X, pB.Y)
		dc.DrawLine(pB.X, pB.Y, pC.X, pC.Y)
		dc.DrawLine(pC.X, pC.Y, p1.X, p1.Y)
		dc.Stroke()
		return
	}

	kochRecursive(dc, p0, pA, maxDepth-1, jitter)
	kochRecursive(dc, pA, pB, maxDepth-1, jitter)
	kochRecursive(dc, pB, pC, maxDepth-1, jitter)
	kochRecursive(dc, pC, p1, maxDepth-1, jitter)
}
