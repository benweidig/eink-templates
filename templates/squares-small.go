package templates

import (
	"math"

	"github.com/fogleman/gg"
)

type Squares struct {
	name      string
	spacing   float64
	lineWidth float64
}

func NewSquares(name string, spacing, lineWidth float64) Squares {
	return Squares{
		name:      name,
		spacing:   spacing,
		lineWidth: lineWidth,
	}
}

func (s Squares) Name() string {
	return s.name
}

func (s Squares) Draw(ctx *gg.Context) {

	screenWidth := float64(ctx.Width())
	screenHeight := float64(ctx.Height())

	// we want at least half spacing as surrounding margin
	minMargin := s.spacing / 2.0

	// Calculate how many total dots are needed
	xSquares := int((screenWidth - 2*minMargin) / s.spacing)
	ySquares := int((screenHeight - 2*minMargin) / s.spacing)

	// Center the dots on the canvas
	marginX := math.Floor((screenWidth - float64(xSquares)*s.spacing) / 2.0)
	marginY := math.Floor((screenHeight - float64(ySquares)*s.spacing) / 2.0)

	// Draw the dots
	for y := marginY; y <= screenHeight-marginY; y += s.spacing {
		ctx.DrawLine(marginX, y, screenWidth-marginX, y)
		for x := marginX; x <= screenWidth-marginX; x += s.spacing {
			ctx.DrawLine(x, marginY, x, screenHeight-marginY)
		}
	}

	// Actual drawing operation
	ctx.SetRGB(0.7, 0.7, 0.7)
	ctx.SetLineWidth(s.lineWidth)
	ctx.Stroke()
}
