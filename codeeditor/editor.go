package codeeditor

import (
	"strings"
	"unicode/utf8"
)

type Editor struct {
	Cursor Cursor
	// TODO: linked list is better here
	Lines []string
	// TODO: history
}

type Cursor struct {
	Row int
	Col int
}

func New(text string) *Editor {
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		lines = []string{""}
	}
	return &Editor{
		Lines: lines,
		Cursor: Cursor{
			Row: 0,
			Col: 0,
		},
	}
}

func (editor *Editor) Text() string {
	return strings.Join(editor.Lines, "\n")
}

func (editor *Editor) Column() int {
	length := editor.currentLineLength()
	if editor.Cursor.Col >= length {
		return length
	}
	return editor.Cursor.Col
}

func (editor *Editor) Row() int {
	return editor.Cursor.Row
}

// NewLine adds new line at cursor point.
// Usually happens on pressing return key.
func (editor *Editor) NewLine() {
	editor.addLine(editor.currentToEndline())
	editor.deleteToEndline()
	editor.MoveToNextLine()
	editor.MoveCursorToStartOfLine()
}

func (editor *Editor) MoveToLeft() {
	if editor.IsBeginningOfFile() {
		return
	}
	if editor.IsBeginningOfLine() {
		editor.MoveToPrevLine()
		editor.MoveCursorToEndOfLine()
		return
	}
	editor.Cursor.Col -= 1
}

func (editor *Editor) MoveToRight(n int) {
	for i := 0; i < n; i++ {
		if editor.IsEndOfFile() {
			return
		}
		if editor.IsEndOfLine() {
			editor.MoveToNextLine()
			editor.MoveCursorToStartOfLine()
			continue
		}
		editor.Cursor.Col += 1
	}
}

func (editor *Editor) MoveToPrevLine() {
	if editor.Cursor.Row == 0 {
		return
	}
	editor.Cursor.Row -= 1
}

func (editor *Editor) MoveToNextLine() {
	if editor.Cursor.Row >= len(editor.Lines)-1 {
		return
	}
	editor.Cursor.Row += 1
}

// FIXME: maybe method modify line
func (editor *Editor) InsertChar(r rune) {
	editor.insertString(string(r))
	editor.MoveToRight(1)
}

func (editor *Editor) BackwardDeleteChar() {
	if editor.IsBeginningOfFile() {
		return
	}

	line := editor.Lines[editor.Cursor.Row]

	// at the beginning of line
	// attach to previous and delete cur line
	if editor.IsBeginningOfLine() {
		editor.deleteLine()
		editor.MoveToPrevLine()
		editor.MoveCursorToEndOfLine()
		editor.insertString(line)
		return
	}
	line = line[:editor.Column()-1] + line[editor.Column():]
	editor.Lines[editor.Cursor.Row] = line
	editor.MoveToLeft()
}

func (editor *Editor) IsBeginningOfLine() bool {
	return editor.Column() == 0
}

func (editor *Editor) IsEndOfLine() bool {
	return editor.Column() == editor.currentLineLength()
}

func (editor *Editor) IsBeginningOfFile() bool {
	return editor.Column() == 0 && editor.Row() == 0
}

func (editor *Editor) IsEndOfFile() bool {
	return editor.Row() == len(editor.Lines)-1 && editor.IsEndOfLine()
}

func (editor *Editor) MoveCursorToStartOfLine() {
	editor.Cursor.Col = 0
}

func (editor *Editor) MoveCursorToEndOfLine() {
	editor.Cursor.Col = editor.currentLineLength()
}

func (editor *Editor) CharAtCursor() string {
	// FIXME: this is strange
	if editor.IsEndOfLine() {
		return "\n"
	}
	return string([]rune(editor.currentLine())[editor.Column()])
}

func (editor *Editor) currentLine() string {
	return editor.Lines[editor.Cursor.Row]
}

func (editor *Editor) currentLineLength() int {
	return utf8.RuneCountInString(editor.Lines[editor.Cursor.Row])
}

func (editor *Editor) currentToEndline() string {
	col := editor.Cursor.Col
	length := editor.currentLineLength()
	if col > length {
		col = length
	}
	return editor.currentLine()[col:length]
}

func (editor *Editor) deleteToEndline() {
	line := editor.currentLine()
	col := editor.Cursor.Col
	if col >= editor.currentLineLength() {
		col = editor.currentLineLength()
	}
	line = line[:col]
	editor.Lines[editor.Cursor.Row] = line
}

func (editor *Editor) deleteLine() {
	result := []string{}
	result = append(result, editor.Lines[:editor.Cursor.Row]...)
	result = append(result, editor.Lines[editor.Cursor.Row+1:]...)
	editor.Lines = result
}

func (editor *Editor) addLine(newLine string) {
	result := []string{}
	result = append(result, editor.Lines[:editor.Cursor.Row+1]...)
	result = append(result, newLine)
	result = append(result, editor.Lines[editor.Cursor.Row+1:]...)
	editor.Lines = result
}

// maybe method modify line
func (editor *Editor) insertString(s string) {
	line := editor.currentLine()
	col := editor.Column()
	line = line[:col] + s + line[col:]
	editor.Lines[editor.Cursor.Row] = line
}
