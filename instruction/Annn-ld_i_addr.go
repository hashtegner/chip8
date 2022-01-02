package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldIAddr struct {
	opcode chip8.Opcode
}

// LDIAddr - Annn - Set I = nnn
func LDIAddr(opcode chip8.Opcode) Instruction {
	return &ldIAddr{opcode: opcode}
}

func (i *ldIAddr) Execute(c *chip8.Chip8) error {
	c.Index = i.opcode.NNN

	c.ProgramCounter += 2
	return nil
}

func (i *ldIAddr) String() string {
	return fmt.Sprintf("LD I addr=0x%x", i.opcode.NNN)
}
