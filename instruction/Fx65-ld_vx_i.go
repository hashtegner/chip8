package instruction

import "github.com/hashtegner/chip8/chip8"

// LDIVx - Fx65 - Read registers V0 through Vx from memory starting at location I
func LDVxI(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
