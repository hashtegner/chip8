package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldIVx struct {
	opcode chip8.Opcode
}

// LDIVx - Fx55 - Store registers V0 through Vx in memory starting at location I
func LDIVx(opcode chip8.Opcode) Instruction {
	return &ldIVx{opcode: opcode}
}

func (i *ldIVx) Execute(c *chip8.Chip8) error {
	var j uint8
	for j = 0; j < uint8(i.opcode.Vx); j++ {
		c.Memory[uint8(c.Index)+j] = c.VX[j]
	}

	c.Index += i.opcode.Vx + 1
	c.ProgramCounter += 2

	return nil
}

func (i *ldIVx) String() string {
	return fmt.Sprintf("LD I Vx=0x%x", i.opcode.Vx)
}
