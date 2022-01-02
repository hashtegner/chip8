package instruction

import "github.com/hashtegner/chip8/chip8"

// SEVXVY - 5xkk - Skip next instruction if Vx == Vy
func SEVXVY(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
