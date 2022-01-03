package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type orVxVy struct {
	opcode chip8.Opcode
}

// OrVxVy - 8xy1 - Set Vx = Vx OR Vy
func OrVxVy(opcode chip8.Opcode) Instruction {
	return &orVxVy{opcode: opcode}
}

func (i *orVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]
	c.VX[i.opcode.Vx] = xValue | yValue

	c.ProgramCounter += 2

	return nil
}

func (i *orVxVy) String() string {
	return fmt.Sprintf("OR Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
