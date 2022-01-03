package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldDtVx struct {
	opcode chip8.Opcode
}

// LDDtVx - Fx15 - Set delay time = Vx
func LDDtVx(opcode chip8.Opcode) Instruction {
	return &ldDtVx{opcode: opcode}
}

func (i *ldDtVx) Execute(c *chip8.Chip8) error {
	c.DelayTimer = c.VX[i.opcode.Vx]
	c.ProgramCounter += 2

	return nil
}

func (i *ldDtVx) String() string {
	return fmt.Sprintf("LD DT Vx=0x%x", i.opcode.Vx)
}
