package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type xorVxVy struct {
	opcode chip8.Opcode
}

// XorVxVy - 8xy3 - Set Vx = Vx XOR Vy
func XorVxVy(opcode chip8.Opcode) Instruction {
	return &xorVxVy{opcode: opcode}
}

func (i *xorVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]
	c.VX[i.opcode.Vx] = xValue ^ yValue

	c.ProgramCounter += 2

	return nil
}

func (i *xorVxVy) String() string {
	return fmt.Sprintf("XOR Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
