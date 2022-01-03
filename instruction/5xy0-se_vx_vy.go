package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type seVxVy struct {
	opcode chip8.Opcode
}

// SEVxVy - 5xkk - Skip next instruction if Vx == Vy
func SEVxVy(opcode chip8.Opcode) Instruction {
	return &seVxVy{opcode: opcode}
}

func (i *seVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]

	if xValue == yValue {
		c.ProgramCounter += 4
	} else {
		c.ProgramCounter += 2
	}

	return nil
}

func (i *seVxVy) String() string {
	return fmt.Sprintf("SE Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
