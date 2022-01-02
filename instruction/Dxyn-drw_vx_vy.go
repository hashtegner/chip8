package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type drwVxVy struct {
	opcode chip8.Opcode
}

// DrnVxVy - Dxyn - Display n-byte sprite strating at memory location I at (Vx, Vy), set VF = collision
func DrwVxVy(opcode chip8.Opcode) Instruction {
	return &drwVxVy{opcode: opcode}
}

func (i *drwVxVy) Execute(c *chip8.Chip8) error {
	colRegister := c.VX[i.opcode.Vx]
	rowRegister := c.VX[i.opcode.Vy]
	height := i.opcode.N

	// set collision flag
	c.VX[0xF] = 0

	var row uint16
	var bitIndex uint8
	for row = 0; row < height; row++ {
		// sprite has 8 bits size
		sprite := c.Memory[c.Index+row]

		for bitIndex = 0; bitIndex < 8; bitIndex++ {
			// the value of bit in the sprite
			bitValue := (sprite >> bitIndex) & 0x1

			rowCoord := (rowRegister + uint8(row)) % chip8.ScreenHeight
			colCoord := (colRegister + uint8(7-bitIndex)) % chip8.ScreenWidth

			currentPixel := c.Screen[rowCoord][colCoord]

			// pixel will colide
			if bitValue == 1 && currentPixel == 1 {
				c.VX[0xF] = 1
			}

			c.Screen[rowCoord][colCoord] = currentPixel ^ bitValue
		}
	}

	c.Draw = true
	c.ProgramCounter += 2

	return nil
}

func (i *drwVxVy) String() string {
	return fmt.Sprintf("DRW Vx=0x%x Vy=0x%x height=0x%x", i.opcode.Vx, i.opcode.Vx, i.opcode.N)
}
