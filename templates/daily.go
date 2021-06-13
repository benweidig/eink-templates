package templates

import (
	"fmt"

	"github.com/fogleman/gg"
)

type Daily struct {
	spacing   float64
	dotSize   float64
	startHour int
	endHour   int
}

func NewDaily(spacing float64, dotSize float64, startHour, endHour int) Daily {
	return Daily{
		spacing:   spacing,
		dotSize:   dotSize,
		startHour: startHour,
		endHour:   endHour,
	}
}

func (d Daily) Name() string {
	return "daily"
}

func (d Daily) Draw(ctx *gg.Context) {

	ctx.SetRGB(0.5, 0.5, 0.5)

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

	// horizontal Line for date / topic
	ctx.DrawLine(startX, startY+1.5*d.spacing, startX+float64(xDotsTotal)*d.spacing, startY+1.5*d.spacing)

	d.drawTimebar(ctx, startX, startY+2.0*d.spacing, yDotsTotal-3)

	// vertical divider between timebar and note area
	ctx.DrawLine(
		startX+float64(int(0.4*float64(xDotsTotal)))*d.spacing,
		startY+2.0*d.spacing,
		startX+float64(int(0.4*float64(xDotsTotal)))*d.spacing,
		startY+float64(yDotsTotal-1)*d.spacing)

	ctx.SetLineWidth(d.dotSize)
	ctx.Stroke()

	// Draw the dots
	for dotsY := 0; dotsY < yDotsTotal; dotsY++ {
		for dotsX := 0; dotsX <= xDotsTotal; dotsX++ {
			x := float64(dotsX)*d.spacing + startX - d.dotSize/2.0
			y := float64(dotsY)*d.spacing + startY - d.dotSize/2.0
			d.drawDot(ctx, x, y)
		}
	}

	ctx.Fill()
}

func (d Daily) drawTimebar(ctx *gg.Context, originX float64, originY float64, verticalSpaces int) {
	// +--+
	// |<-|
	// +--+
	ctx.DrawLine(
		originX,
		originY,
		originX,
		originY+d.spacing*float64(verticalSpaces))

	// +--+
	// |->|
	// +--+
	ctx.DrawLine(
		originX+d.spacing,
		originY,
		originX+d.spacing,
		originY+d.spacing*float64(verticalSpaces))

	// +--+
	// |^^|
	// +--+
	ctx.DrawLine(
		originX,
		originY,
		originX+d.spacing,
		originY)

	// +--+
	// |vv|
	// +--+
	ctx.DrawLine(
		originX,
		originY+d.spacing*float64(verticalSpaces),
		originX+d.spacing,
		originY+d.spacing*float64(verticalSpaces))

	// hour labels and half-hour dots

	ctx.LoadFontFace("./fonts/ComicMono.ttf", 32)

	labelSpacing := (d.spacing * float64(verticalSpaces)) / float64(d.endHour-d.startHour)

	midDotX := originX + 0.5*d.spacing
	midDotY := originY + 0.5*labelSpacing

	d.drawDot(ctx, midDotX, midDotY)

	for timelineY := 1; timelineY < d.endHour-d.startHour; timelineY++ {
		ctx.DrawStringAnchored(
			fmt.Sprintf("%02d", d.startHour+timelineY),
			(originX + 0.5*d.spacing),
			originY+float64(timelineY)*labelSpacing,
			0.5,
			0.5)

		midDotY := originY + (float64(timelineY)+0.5)*labelSpacing
		d.drawDot(ctx, midDotX, midDotY)
	}
}

func (d Daily) drawDot(ctx *gg.Context, x float64, y float64) {
	ctx.DrawRectangle(x, y, d.dotSize, d.dotSize)
}
