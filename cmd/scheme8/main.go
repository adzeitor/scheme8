// graphical version of scheme
package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/adzeitor/goscheme/scheme"
)

const scale = 7

var picoColors = [...]sdl.Color{
	sdl.Color{R: 0, G: 0, B: 0, A: 255},
	sdl.Color{R: 29, G: 43, B: 83, A: 255},
	sdl.Color{R: 126, G: 37, B: 83, A: 255},
	sdl.Color{R: 0, G: 135, B: 81, A: 255},
	sdl.Color{R: 171, G: 82, B: 54, A: 255},
	sdl.Color{R: 95, G: 87, B: 79, A: 255},
	sdl.Color{R: 194, G: 195, B: 199, A: 255},
	sdl.Color{R: 255, G: 241, B: 232, A: 255},
	sdl.Color{R: 255, G: 0, B: 77, A: 255},
}

var buttons = [...]sdl.Scancode{
	// left
	0: 80,
	// right
	1: 79,
	// up
	2: 82,
	// down
	3: 81,
}

func main() {
	window, err := initSDL()
	if err != nil {
		panic(err)
	}
	defer func() {
		window.Destroy()
		sdl.Quit()
	}()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	font, err := ttf.OpenFont("monogram-extended.ttf", 7*scale)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "OpenFont: %s\n", err)
		panic(err)
	}

	running := true

	keys := map[sdl.Scancode]bool{}
	ctx := &EngineContext{
		Surface: surface,
		Keys:    keys,
		Time:    0,
		Font:    font,
		Editor:  true,
	}
	env := environmentWithSDL(ctx)

	//source, err := os.ReadFile("scheme8.scm")
	//if err != nil {
	//	panic(err)
	//}
	//_, env = scheme.EvalBuffer(string(source), env)
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case sdl.QuitEvent:
				running = false
				//err := os.WriteFile("scheme8.scm.bak", []byte(editor.Text()), 0600)
				//if err != nil {
				//	panic(err)
				//}
				//err = os.Rename("scheme8.scm.bak", "scheme8.scm")
				//if err != nil {
				//	panic(err)
				//}
				break
			case sdl.KeyboardEvent:
				ctx.Keys[e.Keysym.Scancode] = e.State == sdl.PRESSED
				if ctx.Keys[sdl.SCANCODE_ESCAPE] {
					ctx.Editor = !ctx.Editor
					if !ctx.Editor {
						_, env = scheme.EvalBuffer(editor.Text(), env)
					}
				}
				if e.State == sdl.PRESSED {
					ctx.Pressed = append(ctx.Pressed, e.Keysym)
				}
			}
		}
		if ctx.Editor {
			editorTime += 1
			updateEditor(ctx)
			drawEditor(ctx)
		} else {
			ctx.Time += 1
			env.Global["time"] = ctx.Time
			_, env = scheme.EvalInEnvironment(`
				(do 
					(_update)
					(_draw)
				)
			`, env)
		}

		window.UpdateSurface()
		// do not trigger released buttons twice
		ctx.Pressed = nil
		// FIXME: DIRTY HACK for 60 fps
		time.Sleep(time.Second / 60)
	}
}

func initSDL() (*sdl.Window, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	ttf.Init()

	window, err := sdl.CreateWindow("scheme8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 128*scale, 128*scale, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	surface, err := window.GetSurface()
	if err != nil {
		return nil, err
	}
	surface.FillRect(nil, 0)

	for i := 0; i < 64; i++ {
		rect := sdl.Rect{
			X: int32(rand.Int()%128) * scale,
			Y: int32(rand.Int()%128) * scale,
			W: scale * int32(i),
			H: scale * int32(i),
		}
		colour := picoColors[i/20%8]
		pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
		surface.FillRect(&rect, pixel)
		window.UpdateSurface()
	}

	return window, nil
}
