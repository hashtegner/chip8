package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldVxVy struct {
	opcode chip8.Opcode
}

// LDVxVy - 8xy0 - Set Vx = Vy
func LDVxVy(opcode chip8.Opcode) Instruction {
	return &ldVxVy{opcode: opcode}
}

func (i *ldVxVy) Execute(c *chip8.Chip8) error {
	value := c.VX[i.opcode.Vy]
	c.VX[i.opcode.Vx] = value

	c.ProgramCounter += 2

	return nil
}

func (i *ldVxVy) String() string {
	return fmt.Sprintf("LD Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
