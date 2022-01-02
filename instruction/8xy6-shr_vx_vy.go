package instruction

import "github.com/hashtegner/chip8/chip8"

// SHRVxVyLeastSignificant - 8xy6 - Set Vx = Vx SHR 1 if the least-significant bit of Vx is 1.
func SHRVxVyLeastSignificant(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
