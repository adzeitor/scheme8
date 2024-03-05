package codeeditor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEditor_NewLine(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		// arrange
		editor := New("")

		// act
		editor.NewLine()

		// assert
		require.Equal(t, "\n", editor.Text())
	})

	t.Run("start of file", func(t *testing.T) {
		// arrange
		editor := New("foo")

		// act
		editor.NewLine()

		// assert
		require.Equal(t, "\nfoo", editor.Text())
	})

	t.Run("between lines", func(t *testing.T) {
		// arrange
		editor := New("foobar")
		// move to 'b'
		editor.MoveToRight(3)

		// act
		editor.NewLine()

		// assert
		require.Equal(t, "foo\nbar", editor.Text())
		require.Equal(t, 0, editor.Column(), "cursor should be at the beginning")
		require.Equal(t, 1, editor.Row(), "cursor should be at the next line")
	})
}

func TestEditor_BackwardDeleteChar(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		// arrange
		editor := New("")

		// act
		editor.BackwardDeleteChar()

		// assert
		require.Equal(t, "", editor.Text())
	})

	t.Run("cursor at start of file (do nothing)", func(t *testing.T) {
		// arrange
		editor := New("foo")

		// act
		editor.BackwardDeleteChar()

		// assert
		require.Equal(t, "foo", editor.Text())
		require.Equal(t, 0, editor.Column(), "cursor should stay at beginning")
	})

	t.Run("delete chars", func(t *testing.T) {
		// arrange
		editor := New("foobar")
		// go to 'b'
		editor.MoveToRight(3)

		// act
		editor.BackwardDeleteChar()
		editor.BackwardDeleteChar()

		// assert
		require.Equal(t, "fbar", editor.Text())
		// cursor stays at b
		require.Equal(t, "b", editor.CharAtCursor())
	})

	t.Run("join lines if delete at beginning of line", func(t *testing.T) {
		// arrange
		editor := New("foo\nbar")
		// go to 'b'
		editor.MoveToNextLine()

		// act
		editor.BackwardDeleteChar()

		// assert
		require.Equal(t, "foobar", editor.Text())
		// cursor stays at b
		require.Equal(t, "b", editor.CharAtCursor())
	})

	t.Run("should delete new line", func(t *testing.T) {
		// arrange
		editor := New("foobar\n")
		// go to '\n'
		editor.MoveToNextLine()

		// act
		editor.BackwardDeleteChar()

		// assert
		require.Equal(t, "foobar", editor.Text())
	})
}

func TestEditor_MoveToLeft(t *testing.T) {
	t.Run("do nothing at beginning of first line", func(t *testing.T) {
		// arrange
		editor := New("foobar")

		// act
		editor.MoveToLeft()

		// assert
		// cursor stays at f
		require.Equal(t, "f", editor.CharAtCursor())
	})

	t.Run("move to previous line at beginning of line", func(t *testing.T) {
		// arrange
		editor := New("foo\nbar")
		editor.MoveToNextLine()

		// act
		editor.MoveToLeft()

		// assert
		require.Equal(t, "\n", editor.CharAtCursor())
	})
}

func TestEditor_MoveToRight(t *testing.T) {
	t.Run("do nothing at the end of file", func(t *testing.T) {
		// arrange
		editor := New("foobar")

		// act
		editor.MoveToRight(7)

		// assert
		require.Equal(t, "\n", editor.CharAtCursor())
	})

	t.Run("move to previous line at beginning of line", func(t *testing.T) {
		// arrange
		editor := New("\nfoobar")

		// act
		editor.MoveToRight(1)

		// assert
		require.Equal(t, "f", editor.CharAtCursor())
	})
}

func TestEditor_MoveToNextLine(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		// arrange
		editor := New("")

		// act
		editor.MoveToNextLine()

		// assert
		// cursor stays at f
		require.Equal(t, "\n", editor.CharAtCursor())
	})
}
