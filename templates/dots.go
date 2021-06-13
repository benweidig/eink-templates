package templates

import (
	"github.com/fogleman/gg"
)

type Dots struct {
	spacing float64
}

func NewDots(spacing float64) Dots {
	return Dots{
		spacing: spacing,
	}
}

func (d Dots) Name() string {
	return "dots"
}

func (d Dots) Draw(ctx *gg.Context) {

	dotSize := 4.0

	screenWidth := float64(ctx.Width())
	screenHeight := float64(ctx.Height())

	// we want at least half spacing as surrounding margin
	minMargin := d.spacing / 2.0

	// Calculate how many total dots are needed
	xDotsTotal := int((screenWidth-2*minMargin)/d.spacing) + 1
	yDotsTotal := int((screenHeight - 2*minMargin) / d.spacing)

	// Center the dots on the canvas
	startX := (screenWidth - float64(xDotsTotal-1)*d.spacing) / 2.0
	startY := (screenHeight - float64(yDotsTotal-1)*d.spacing) / 2.0

	// Draw the dots
	for dotsY := 0; dotsY < yDotsTotal; dotsY++ {
		for dotsX := 0; dotsX <= xDotsTotal; dotsX++ {
			x := float64(dotsX)*d.spacing + startX - dotSize/2.0
			y := float64(dotsY)*d.spacing + startY - dotSize/2.0
			ctx.DrawRectangle(x, y, dotSize, dotSize)
		}
	}

	// Actual drawing operation
	ctx.SetRGB(0.5, 0.5, 0.5)
	ctx.Fill()
}
