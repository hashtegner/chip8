package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type subnVxVy struct {
	opcode chip8.Opcode
}

// SubnVxVy - 8xy7 - Set Vx = Vy - Vx, set VF = Not borrow.
func SubnVxVy(opcode chip8.Opcode) Instruction {
	return &subnVxVy{opcode: opcode}
}

func (i *subnVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]

	c.VX[i.opcode.Vx] = yValue - xValue
	if yValue > xValue {
		c.VX[0xF] = 1
	} else {
		c.VX[0xF] = 0
	}

	c.ProgramCounter += 2

	return nil
}

func (i *subnVxVy) String() string {
	return fmt.Sprintf("SUBN Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
