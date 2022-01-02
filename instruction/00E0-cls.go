package instruction

import (
	"github.com/hashtegner/chip8/chip8"
)

type cls struct {
}

// Cls - 00E0 - Clear the display
func Cls(opcode chip8.Opcode) Instruction {
	return &cls{}
}

func (i *cls) Execute(c *chip8.Chip8) error {
	for row := 0; row < chip8.ScreenHeight; row++ {
		for col := 0; col < chip8.ScreenWidth; col++ {
			c.Screen[row][col] = 0x0
		}
	}

	c.Draw = true
	c.ProgramCounter += 2

	return nil
}

func (i *cls) String() string {
	return "CLS"
}
