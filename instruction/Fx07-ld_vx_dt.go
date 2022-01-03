package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldVxDt struct {
	opcode chip8.Opcode
}

// LDVxDt - Fx07 - Set Vx = delay time value
func LDVxDt(opcode chip8.Opcode) Instruction {
	return &ldVxDt{opcode: opcode}
}

func (i *ldVxDt) Execute(c *chip8.Chip8) error {
	c.VX[i.opcode.Vx] = c.DelayTimer
	c.ProgramCounter += 2

	return nil
}

func (i *ldVxDt) String() string {
	return fmt.Sprintf("LD Vx=0x%x DT", i.opcode.Vx)
}
