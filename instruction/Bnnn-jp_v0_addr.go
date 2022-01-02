package instruction

import "github.com/hashtegner/chip8/chip8"

// JpV0Addr - Bnnn - Jump to location nnn + v0
func JpV0Addr(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
