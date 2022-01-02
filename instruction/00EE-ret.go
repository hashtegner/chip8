package instruction

import "github.com/hashtegner/chip8/chip8"

type ret struct{}

// Ret - 00EE - Return from a subroutine
func Ret(opcode chip8.Opcode) Instruction {
	return &ret{}
}

func (i *ret) Execute(c *chip8.Chip8) error {
	c.StackPointer -= 1
	c.ProgramCounter = c.Stack[c.StackPointer]
	return nil
}

func (i *ret) String() string {
	return "RET"
}
