package main

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/adzeitor/scheme8/codeeditor"
)

var editor = codeeditor.New(`
(define x 0)
(define y 0)
(define x1 30)
(define y1 20)
(define dx 1)
(define dy 1)

(define _update (lambda (prev_state)
  (do
    (if (btn 0) (set! x (- x 1)) #f)
    (if (btn 1) (set! x (+ x 1)) #f)
    (if (btn 2) (set! y (- y 1)) #f)
    (if (btn 3) (set! y (+ y 1)) #f)
    (set! x1 (+ x1 dx))
    (set! y1 (+ y1 dy))
    (if (> x1 100) (set! dx -1) #f)
    (if (> y1 100) (set! dy -1) #f)
    (if (< x1 0) (set! dx 1) #f)
    (if (< y1 0) (set! dy 1) #f)
)))

(define _draw (lambda ()
  (do
    (cls 0)
    (rectfill 
          x y 
          18 18 
          2)
    (rectfill 
          (+ x 1) (+ y 1) 
          17 17
          8)
    (print "hello world" x1 y1)
  )))
`)

// FIXME: get rid of global variables and make structure
var (
	scrollRow  = 0
	scrollCol  = 0
	editorTime = 0
)

func drawEditor(ctx *EngineContext) {
	x := 0
	y := 16
	cls(ctx, 5)
	if (editorTime/15)%2 == 0 {
		rectFill(ctx, x+(editor.Column())*4, y+(editor.Row()-scrollRow)*4, 4, 4, 8)
	}
	for _, line := range editor.Lines[scrollRow:] {
		printText(ctx, line, x, y, 7)
		y += 4
	}
	rectFill(ctx, 0, 122, 128, 6, 8)
	printText(ctx, fmt.Sprintf("LINE %d/%d", editor.Row()+1, len(editor.Lines)), 0, 123, 0)
}

func updateEditor(ctx *EngineContext) {
	if ctx.Pressed == nil {
		return
	}

	for _, keysym := range ctx.Pressed {
		switch keysym.Scancode {
		case sdl.SCANCODE_UP:
			editor.MoveToPrevLine()
		case sdl.SCANCODE_DOWN:
			editor.MoveToNextLine()
		case sdl.SCANCODE_LEFT:
			editor.MoveToLeft()
		case sdl.SCANCODE_RIGHT:
			editor.MoveToRight(1)
		case sdl.SCANCODE_BACKSPACE:
			editor.BackwardDeleteChar()
		case sdl.SCANCODE_RETURN:
			editor.NewLine()
		default:
			// FIXME: so many hacks here
			// should be easy way to do it
			s := sdl.GetScancodeName(keysym.Scancode)
			// 9 and shift => '('
			if s == "9" && keysym.Mod&1 >= 1 {
				editor.InsertChar('(')
				break
			}
			// 0 and shift => ')'
			if s == "0" && keysym.Mod&1 >= 1 {
				editor.InsertChar(')')
				break
			}
			if len(s) == 1 {
				editor.InsertChar(rune(strings.ToLower(s)[0]))
				break
			}
			if s == "Space" {
				editor.InsertChar(' ')
				break
			}
		}
	}

	if editor.Row()-scrollRow > 20 {
		scrollRow += 1
	}
	if editor.Row()-scrollRow < 0 {
		scrollRow -= 1
	}
	if scrollRow >= editor.Row() {
		scrollRow = 0
	}

	if editor.Column()-scrollCol > 20 {
		scrollCol += 1
	}
	if editor.Column()-scrollCol < 0 {
		scrollCol -= 1
	}
	if scrollCol >= editor.Column() {
		scrollCol = 0
	}
}
