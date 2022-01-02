package instruction

import "github.com/hashtegner/chip8/chip8"

// LDDtVx - Fx15 - Set delay time = Vx
func LDDtVx(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
