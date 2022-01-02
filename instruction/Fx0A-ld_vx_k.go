package instruction

import "github.com/hashtegner/chip8/chip8"

// LDVxK - Fx0A - Wait for key press, store the value of the key in Vx
func LDVxK(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
