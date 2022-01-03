package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldfVx struct {
	opcode chip8.Opcode
}

// LDFVx - Fx29 - Set I = location of sprite for digit Vx
func LDFVx(opcode chip8.Opcode) Instruction {
	return &ldfVx{opcode: opcode}
}

func (i *ldfVx) Execute(c *chip8.Chip8) error {
	c.Index = uint16(c.VX[i.opcode.Vx])
	c.ProgramCounter += 2

	return nil
}

func (i *ldfVx) String() string {
	return fmt.Sprintf("LD F vx=0x%x", i.opcode.Vx)
}
