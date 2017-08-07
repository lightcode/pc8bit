package main

func main() {
	mem := new(Memory)
	mem.Write(0x00, (I_LDA<<4)|0x0E)
	mem.Write(0x01, (I_ADD<<4)|0x0F)
	mem.Write(0x03, (I_ADD<<4)|0x0D)
	mem.Write(0x04, (I_OUT<<4)|0x00)
	mem.Write(0x05, (I_HLT<<4)|0x00)
	mem.Write(0x0D, 10)
	mem.Write(0x0E, 15)
	mem.Write(0x0F, 27)
	cpu := &CPU{memory: mem}
	cpu.Reset()
	cpu.Run()
}
