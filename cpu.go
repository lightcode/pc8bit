package main

import "fmt"
import "time"
import "github.com/fatih/color"

const INSTRUCTION_LENGTH = 8

type CPU struct {
	clockEnabled   bool    // Is our clock enabled?
	pc             byte    // Program counter
	regA           byte    // A register
	regB           byte    // B register
	instructionReg byte    // Instruction register (4 bits)
	cycle          byte    // Operation Counter
	bus            byte    // Bus
	memAddrReg     byte    // Memory Address Register
	memory         *Memory // 16 bytes memory
}

func (cpu *CPU) Reset() {
	cpu.pc = 0
	cpu.regA = 0
	cpu.regB = 0
	cpu.clockEnabled = true
}

func (cpu *CPU) Run() {
	for cpu.clockEnabled {
		cpu.tick()
		time.Sleep(20 * time.Millisecond)
	}
}

func (cpu *CPU) tick() {
	fmt.Printf("PC = %d, cycle = %d\n", cpu.pc, cpu.cycle)
	fmt.Printf("Bus = %08b\n", cpu.bus)
	fmt.Printf("RegIns = %08b   MemReg = %04b\n", cpu.instructionReg, cpu.memAddrReg)
	fmt.Printf("RegA   = %08b   RegB   = %08b\n", cpu.regA, cpu.regB)

	// Loading instruction into the instruction register
	var ins byte = (cpu.instructionReg >> 4) & 0xF
	var moff byte = ins*INSTRUCTION_LENGTH + cpu.cycle
	cpu.runMicrocode(microcode[moff])

	// Increment op counter
	cpu.cycle = (cpu.cycle + 1) % 8

	if cpu.cycle == 0 {
		fmt.Println("========================================")
	} else {
		fmt.Println("----------------------------------------")
	}
}

func (cpu *CPU) runMicrocode(op Opcode) {
	for _, m := range mins {
		if (op & m) > 0 {
			cpu.runMicroInstruction(m)
		}
	}
}

func (cpu *CPU) runMicroInstruction(inst Opcode) {
	switch inst {
	case C_CI:
		color.Red("C_CI")
		cpu.pc++
	case C_CO:
		color.Red("C_CO")
		cpu.bus = cpu.pc
	case C_RO:
		color.Red("C_RO")
		cpu.bus = cpu.memory.Read(cpu.memAddrReg)
	case C_IO:
		color.Red("C_IO")
		cpu.bus = cpu.instructionReg & 0x0F
	case C_MI:
		color.Red("C_MI")
		cpu.memAddrReg = cpu.bus & 0x0F
	case C_II:
		color.Red("C_II")
		cpu.instructionReg = cpu.bus
	case C_HALT:
		color.Red("HALT")
		cpu.clockEnabled = false
	case C_OI:
		color.Red("OI")
		color.Cyan("OUT : %08b (%d)", cpu.bus, cpu.bus)
	case C_AO:
		color.Red("AO")
		cpu.bus = cpu.regA
	case C_AI:
		color.Red("AI")
		cpu.regA = cpu.bus
	case C_BO:
		color.Red("BO")
		cpu.bus = cpu.regB
	case C_BI:
		color.Red("BI")
		cpu.regB = cpu.bus
	case C_EO:
		color.Red("EO")
		cpu.bus = (cpu.regA + cpu.regB) & 0xFF
	}
}
