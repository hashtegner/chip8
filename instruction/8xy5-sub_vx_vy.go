package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type subVxVy struct {
	opcode chip8.Opcode
}

// SubVxVy - 8xy5 - Set Vx = Vx - Vy, set VF = Not borrow
func SubVxVy(opcode chip8.Opcode) Instruction {
	return &subVxVy{opcode: opcode}
}

func (i *subVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]
	c.VX[i.opcode.Vx] = xValue - yValue

	if xValue > yValue {
		c.VX[0xF] = 1
	} else {
		c.VX[0xF] = 0
	}

	c.ProgramCounter += 2

	return nil
}

func (i *subVxVy) String() string {
	return fmt.Sprintf("SUB Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
