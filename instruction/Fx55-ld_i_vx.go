package instruction

import "github.com/hashtegner/chip8/chip8"

// LDIVx - Fx55 - Store registers V0 through Vx in memory starting at location I
func LDIVx(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
