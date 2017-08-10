package main

func main() {
	mem := new(Memory)
	mem.Write(0x00, I_LDA)
	mem.Write(0x01, 0xFE)
	mem.Write(0x02, I_ADD)
	mem.Write(0x03, 0xFF)
	mem.Write(0x04, I_ADD)
	mem.Write(0x05, 0xFD)
	mem.Write(0x06, I_OUT)
	mem.Write(0x07, 0x00)
	//mem.Write(0x08, I_JMP<<4)
	//mem.Write(0x09, 0x04)
	mem.Write(0x0A, I_HLT)
	mem.Write(0x0B, 0x00)
	mem.Write(0xFD, 10)
	mem.Write(0xFE, 15)
	mem.Write(0xFF, 27)
	cpu := &CPU{memory: mem}
	cpu.Reset()
	cpu.Run()
}
