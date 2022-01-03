package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldVxI struct {
	opcode chip8.Opcode
}

// LDVxI - Fx65 - Read registers V0 through Vx from memory starting at location I
func LDVxI(opcode chip8.Opcode) Instruction {
	return &ldVxI{opcode: opcode}
}

func (i *ldVxI) Execute(c *chip8.Chip8) error {
	var j uint16
	for j = 0; j < i.opcode.Vx; j++ {
		c.Memory[c.Index+j] = c.VX[j]
	}

	c.Index += i.opcode.Vx + 1
	c.ProgramCounter += 2

	return nil
}

func (i *ldVxI) String() string {
	return fmt.Sprintf("LD vx=0x%x I", i.opcode.Vx)
}
