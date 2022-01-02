package instruction

import "github.com/hashtegner/chip8/chip8"

// LDFVx - Fx29 - Set I = location of sprite for digit Vx
func LDFVx(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
