package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type addVxByte struct {
	opcode chip8.Opcode
}

// AddVxByte - 7xkk - Set Vx = Vx + kk
func AddVxByte(opcode chip8.Opcode) Instruction {
	return &addVxByte{opcode: opcode}
}

func (i *addVxByte) Execute(c *chip8.Chip8) error {
	c.VX[i.opcode.Vx] += uint8(i.opcode.KK)

	c.ProgramCounter += 2

	return nil
}

func (i *addVxByte) String() string {
	return fmt.Sprintf("ADD vx=0x%x byte=0x%x", i.opcode.Vx, i.opcode.KK)
}
