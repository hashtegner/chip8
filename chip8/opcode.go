package chip8

import "fmt"

type Opcode struct {
	Raw     uint16 // original opcode
	Decoded uint16 // the highest order byte
	Vx      uint16 // the lower 4 bits of the high byte
	Vy      uint16 // the upper 4 bits of the low byte
	N       uint16 // the lowest 4 bits
	KK      uint16 // the lowest 4 bits
	NNN     uint16 // the lowest 12 bits
}

func (o Opcode) String() string {
	return fmt.Sprintf("raw=0x%x decoded=0x%x vx=0x%x vy=0x%x n=0x%x kk=0x%x nnn=0x%x",
		o.Raw,
		o.Decoded,
		o.Vx,
		o.Vy,
		o.N,
		o.KK,
		o.NNN,
	)
}
