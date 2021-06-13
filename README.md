# eInk Templates

This repo is an attempt to generate templates for eInk-tables with code instead of a vector graphics tool (e.g. Inkscape/Illustrator).

Eversything is kind of hardcoded right now, and only 2 different varieties of templates are available.

## Creating Template

Implement the interface `templates.Template`, add it to the logic in `main.go`.


```golang
go run main.go
```

## License

The code itself is CC0, see LICENSE.

Fonts may be licensed differently, see fonts/LICENSE.md for their exact license.
