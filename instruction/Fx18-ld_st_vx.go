package instruction

import "github.com/hashtegner/chip8/chip8"

// LDStVx - Fx18 - Set sound timer = Vx
func LDStVx(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
