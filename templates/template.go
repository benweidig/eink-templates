package templates

import "github.com/fogleman/gg"

type Template interface {
	Name() string

	Draw(ctx *gg.Context)
}
