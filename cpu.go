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
	cpu.printStates()

	// Loading instruction into the instruction register
	var ins byte = (cpu.instructionReg >> 4) & 0xF
	var moff byte = ins*INSTRUCTION_LENGTH + cpu.cycle
	op := microcode[moff]
	for _, m := range mins {
		if (op & m) > 0 {
			cpu.runMicroInstruction(m)
		}
	}

	// Increment op counter
	cpu.cycle = (cpu.cycle + 1) % 8

	if cpu.cycle == 0 {
		color.HiBlack("========================================")
	} else {
		color.HiBlack("----------------------------------------")
	}
}

func (cpu *CPU) runMicroInstruction(inst Opcode) {
	switch inst {
	case C_CI:
		printOpcode("CI")
		cpu.pc++
	case C_CO:
		printOpcode("CO")
		cpu.bus = cpu.pc
	case C_J:
		printOpcode("J")
		cpu.pc = cpu.bus & 0x0F
	case C_RO:
		printOpcode("RO")
		cpu.bus = cpu.memory.Read(cpu.memAddrReg)
	case C_IO:
		printOpcode("IO")
		cpu.bus = cpu.instructionReg & 0x0F
	case C_MI:
		printOpcode("MI")
		cpu.memAddrReg = cpu.bus & 0x0F
	case C_II:
		printOpcode("II")
		cpu.instructionReg = cpu.bus
	case C_HALT:
		printOpcode("HALT")
		cpu.clockEnabled = false
	case C_OI:
		printOpcode("OI")
		color.Cyan("OUT : %08b (%d)", cpu.bus, cpu.bus)
	case C_AO:
		printOpcode("AO")
		cpu.bus = cpu.regA
	case C_AI:
		printOpcode("AI")
		cpu.regA = cpu.bus
	case C_BO:
		printOpcode("BO")
		cpu.bus = cpu.regB
	case C_BI:
		printOpcode("BI")
		cpu.regB = cpu.bus
	case C_EO:
		printOpcode("EO")
		cpu.bus = (cpu.regA + cpu.regB) & 0xFF
	}
}

func printOpcode(text string) {
	color.Red(text)
}

func (cpu *CPU) printStates() {
	fmt.Printf("PC = %d, cycle = %d\n", cpu.pc, cpu.cycle)
	fmt.Printf("Bus = %08b\n", cpu.bus)
	fmt.Printf("RegIns = %08b   MemReg = %04b\n", cpu.instructionReg, cpu.memAddrReg)
	fmt.Printf("RegA   = %08b   RegB   = %08b\n", cpu.regA, cpu.regB)
}
