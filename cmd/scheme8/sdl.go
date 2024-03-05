package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func cls(ctx *EngineContext, color int) {
	rect := sdl.Rect{
		0,
		0,
		128 * scale,
		128 * scale,
	}
	colour := picoColors[color%16]
	pixel := sdl.MapRGBA(ctx.Surface.Format, colour.R, colour.G, colour.B, colour.A)
	ctx.Surface.FillRect(&rect, pixel)
}

func pset(ctx *EngineContext, x int, y int, color int) {
	rectFill(ctx, x, y, scale, scale, color)
}

func rectFill(ctx *EngineContext, x int, y int, w int, h int, c int) {
	rect := sdl.Rect{
		int32(x * scale),
		int32(y * scale),
		int32(w * scale),
		int32(h * scale),
	}
	colour := picoColors[c%16]
	pixel := sdl.MapRGBA(ctx.Surface.Format, colour.R, colour.G, colour.B, colour.A)
	ctx.Surface.FillRect(&rect, pixel)
}

func circFill(ctx *EngineContext, x int, y int, w int, h int, c int) {
	rect := sdl.Rect{
		int32(x * scale),
		int32(y * scale),
		int32(w * scale),
		int32(h * scale),
	}
	colour := picoColors[c%16]
	pixel := sdl.MapRGBA(ctx.Surface.Format, colour.R, colour.G, colour.B, colour.A)
	ctx.Surface.FillRect(&rect, pixel)
}

func printText(ctx *EngineContext, s string, x int, y int, color int) {
	for _, r := range s {
		rendered, _ := ctx.Font.RenderGlyphSolid(r, picoColors[color%16])
		rendered.Blit(nil, ctx.Surface, &sdl.Rect{
			X: int32(x)*scale + scale,
			Y: int32(y)*scale - scale - 2,
			W: 8 * scale,
			H: 8 * scale,
		})
		rendered.Free()
		x += 4
	}
}
