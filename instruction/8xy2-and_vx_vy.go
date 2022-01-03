package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type andVxVy struct {
	opcode chip8.Opcode
}

// AndVxVy - 8xy2 - Set Vx = Vx AND Vy
func AndVxVy(opcode chip8.Opcode) Instruction {
	return &andVxVy{opcode: opcode}
}

func (i *andVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]
	c.VX[i.opcode.Vx] = xValue & yValue

	c.ProgramCounter += 2

	return nil
}

func (i *andVxVy) String() string {
	return fmt.Sprintf("AND Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
