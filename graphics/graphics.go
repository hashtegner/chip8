package graphics

import (
	"github.com/hashtegner/chip8/chip8"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

const screenDelta = 10

type Graphics struct {
	window  *sdl.Window
	surface *sdl.Surface
}

func New(title string, width, height int32) (*Graphics, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, errors.Wrap(err, "error on start std2")
	}

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		width*screenDelta,
		height*screenDelta,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		return nil, errors.Wrap(err, "error on create window")
	}

	surface, err := window.GetSurface()
	if err != nil {
		return nil, errors.Wrap(err, "error on get surface from window")
	}

	return &Graphics{window, surface}, nil
}

func (g *Graphics) Draw(buffer chip8.ScreenBuffer) error {
	surface := g.surface

	surface.FillRect(nil, 0xffffff)

	for rowIndex := 0; rowIndex < len(buffer); rowIndex++ {
		row := buffer[rowIndex]

		for colIndex := 0; colIndex < len(row); colIndex++ {
			var color uint32

			if buffer[rowIndex][colIndex] > 0 {
				color = 0xffffff
			} else {
				color = 0x0
			}

			if err := g.paintCell(rowIndex, colIndex, color); err != nil {
				return errors.Wrap(err, "error on paint")
			}
		}
	}

	return errors.Wrap(g.window.UpdateSurface(), "error on update surface")
}

func (g *Graphics) paintCell(row, col int, color uint32) error {
	rect := &sdl.Rect{
		X: int32(col) * screenDelta,
		Y: int32(row) * screenDelta,
		W: screenDelta,
		H: screenDelta,
	}

	g.surface.FillRect(rect, color)

	return nil
}

func (g *Graphics) Close() {
	g.window.Destroy()
	sdl.Quit()
}
