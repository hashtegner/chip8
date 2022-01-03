package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type sneVxVy struct {
	opcode chip8.Opcode
}

// SNEVxVy - 9xy0 - Skip next instruction if Vx != Vy
func SNEVxVy(opcode chip8.Opcode) Instruction {
	return &sneVxVy{opcode: opcode}
}

func (i *sneVxVy) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	yValue := c.VX[i.opcode.Vy]

	if xValue != yValue {
		c.ProgramCounter += 4
	} else {
		c.ProgramCounter += 2
	}

	return nil
}

func (i *sneVxVy) String() string {
	return fmt.Sprintf("SNE Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
