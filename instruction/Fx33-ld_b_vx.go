package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldbVx struct {
	opcode chip8.Opcode
}

// LDBVx - Fx33 - Store BCD representation of Vx in memory locations I, I+1 and I+2
func LDBVx(opcode chip8.Opcode) Instruction {
	return &ldbVx{opcode: opcode}
}

func (i *ldbVx) Execute(c *chip8.Chip8) error {
	value := c.VX[i.opcode.Vx]

	c.Memory[c.Index] = value / 100          // hundred's digit
	c.Memory[c.Index+1] = (value / 10) % 10  // ten's digit
	c.Memory[c.Index+2] = (value % 100) / 10 // one's digit

	c.ProgramCounter += 2

	return nil
}

func (i *ldbVx) String() string {
	return fmt.Sprintf("LD bd vx=0x%x", i.opcode.Vx)
}
