package instruction

import "github.com/hashtegner/chip8/chip8"

// SkpnVx - Ex9E - Skip next instruction if key with the value of Vx is pressed.
func SkpVx(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
