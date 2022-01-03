package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldStVx struct {
	opcode chip8.Opcode
}

// LDStVx - Fx18 - Set sound timer = Vx
func LDStVx(opcode chip8.Opcode) Instruction {
	return &ldStVx{opcode: opcode}
}

func (i *ldStVx) Execute(c *chip8.Chip8) error {
	c.SoundTimer = c.VX[i.opcode.Vx]
	c.ProgramCounter += 2

	return nil
}

func (i *ldStVx) String() string {
	return fmt.Sprintf("LD St Vx=0x%x", i.opcode.Vx)
}
