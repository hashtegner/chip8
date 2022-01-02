package opcode

import (
	"log"

	"github.com/hashtegner/chip8/chip8"
	"github.com/hashtegner/chip8/instruction"
)

type instructionDecoder = func(opcode chip8.Opcode) instruction.Instruction

var opcodes0000Table = map[uint16]instructionDecoder{
	0x0000: instruction.Cls,
	0x000E: instruction.Ret,
}

var opcodes8000Table = map[uint16]instructionDecoder{
	0x0000: instruction.LDVxVy,
	0x0001: instruction.OrVxVy,
	0x0002: instruction.AndVxVy,
	0x0003: instruction.XorVxVy,
	0x0004: instruction.AddVxVy,
	0x0005: instruction.SubVxVy,
	0x0006: instruction.SHRVxVyLeastSignificant,
	0x0007: instruction.SubnVxVy,
	0x000E: instruction.SHRVxVyMostSignificant,
}

var opcodesE000Table = map[uint16]instructionDecoder{
	0x009E: instruction.SkpVx,
	0x00A1: instruction.SkpnVx,
}

var opcodesF000Table = map[uint16]instructionDecoder{
	0x0007: instruction.LDVxDt,
	0x000A: instruction.LDVxK,
	0x0015: instruction.LDDtVx,
	0x0018: instruction.LDStVx,
	0x001E: instruction.AddIVx,
	0x0029: instruction.LDFVx,
	0x0033: instruction.LDBVx,
	0x0055: instruction.LDIVx,
	0x0065: instruction.LDVxI,
}

var opcodesTable = map[uint16]instructionDecoder{
	0x0000: func(opcode chip8.Opcode) instruction.Instruction {
		return decodeFromTable(opcodes0000Table, opcode, opcode.N)
	},
	0x1000: instruction.JPAddr,
	0x2000: instruction.CallAddr,
	0x3000: instruction.SEVxByte,
	0x4000: instruction.SENVXByte,
	0x5000: instruction.SEVXVY,
	0x6000: instruction.LDVxByte,
	0x7000: instruction.AddVxByte,
	0x8000: func(opcode chip8.Opcode) instruction.Instruction {
		return decodeFromTable(opcodes8000Table, opcode, opcode.N)
	},
	0x9000: instruction.SNEVxVy,
	0xA000: instruction.LDIAddr,
	0xB000: instruction.JpV0Addr,
	0xC000: instruction.RndVxByte,
	0xD000: instruction.DrwVxVy,
	0xE000: func(opcode chip8.Opcode) instruction.Instruction {
		return decodeFromTable(opcodesE000Table, opcode, opcode.KK)
	},
	0xF000: func(opcode chip8.Opcode) instruction.Instruction {
		return decodeFromTable(opcodesF000Table, opcode, opcode.KK)
	},
}

// Fetch opcode from memory and program counter
func Fetch(c *chip8.Chip8) chip8.Opcode {
	a := uint16(c.Memory[c.ProgramCounter]) << 8
	b := uint16(c.Memory[c.ProgramCounter+1])
	opcode := a | b

	return chip8.Opcode{
		Raw:     opcode,
		Decoded: opcode & 0xF000,
		Vx:      (opcode >> 8) & 0x000F,
		Vy:      (opcode >> 4) & 0x000F,
		N:       opcode & 0x000F,
		KK:      opcode & 0x00FF,
		NNN:     opcode & 0x0FFF,
	}
}

// Decode an opcode instruction
func Decode(opcode chip8.Opcode) instruction.Instruction {
	inst := decodeFromTable(opcodesTable, opcode, opcode.Decoded)

	log.Printf("%s", inst.String())

	return inst
}

// decodeFromTable decode opcode from given table and return struction
func decodeFromTable(table map[uint16]instructionDecoder, opcode chip8.Opcode, decoded uint16) instruction.Instruction {
	decode, hasDecoder := table[decoded]

	if !hasDecoder {
		return instruction.Undefined(opcode)
	}

	return decode(opcode)
}
