package main

import (
	"fmt"

	"github.com/benweidig/eink-templates/templates"
	"github.com/fogleman/gg"
)

const (
	screenWidth  = 1404.0
	screenHeight = 1872.0

	outPath = "out"
)

func main() {

	dots := templates.NewDots(48.0)
	squaresSmall := templates.NewSquares("squares-small", 50.0, 4.0)
	squaresLarge := templates.NewSquares("squares-large", 75.0, 4.0)
	daily := templates.NewDaily(48.0, 4.0, 6, 22)

	generate(dots)
	generate(squaresSmall)
	generate(squaresLarge)
	generate(daily)

}

func generate(t templates.Template) {
	ctx := gg.NewContext(screenWidth, screenHeight)

	t.Draw(ctx)

	path := fmt.Sprintf("%s/%s.png", outPath, t.Name())
	ctx.SavePNG(path)
}
