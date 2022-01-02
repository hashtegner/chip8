package instruction

import "github.com/hashtegner/chip8/chip8"

// CallAddr - 2nnn - Call subroutine at nnn
func CallAddr(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
