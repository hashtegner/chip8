package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type undefined struct {
	opcode chip8.Opcode
}

// Undefined - opcode already not implemented
func Undefined(opcode chip8.Opcode) Instruction {
	return &undefined{opcode: opcode}
}

func (i *undefined) Execute(c *chip8.Chip8) error {
	return fmt.Errorf("instruction for opcode %s undefined", i.opcode.String())
}

func (i *undefined) String() string {
	return "UNDEFINED"
}
