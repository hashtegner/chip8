package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type addVxVy struct {
	opcode chip8.Opcode
}

// AddVxVy - 8xy4 - Set Vx = Vx + Vy, set VF = carry
func AddVxVy(opcode chip8.Opcode) Instruction {
	return &addVxVy{opcode: opcode}
}

func (i *addVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]
	value := xValue + yValue

	c.VX[i.opcode.Vx] = value

	if value > 255 {
		c.VX[0xF] = 1
	} else {
		c.VX[0xF] = 0
	}

	c.ProgramCounter += 2

	return nil
}

func (i *addVxVy) String() string {
	return fmt.Sprintf("ADD Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
