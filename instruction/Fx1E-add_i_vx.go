package instruction

import "github.com/hashtegner/chip8/chip8"

// AddIVx - Fx1E - Set I = I + Vx
func AddIVx(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
