package main

import "fmt"
import "time"
import "github.com/fatih/color"

const INSTRUCTION_LENGTH = 8

const (
	_      = iota
	C_CI   = 1 << iota // program Counter Increment
	C_CO               // program Counter Out
	C_RO               // RAM Out
	C_II               // Instruction register In
	C_IO               // Instruction register Out
	C_EO               // ALU Out
	C_AI               // Register A In
	C_AO               // Register A Out
	C_BI               // Register B In
	C_BO               // Register B Out
	C_HALT             // Halt: disable the clock
	C_MI               // Memory Register In
	C_OI               // Out In: bus -> out
)

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
		time.Sleep(200 * time.Millisecond)
	}
}

func (cpu *CPU) tick() {
	fmt.Println("---------------------------")
	fmt.Println("PC = ", cpu.pc)
	fmt.Println("bus = ", cpu.bus)
	fmt.Println("cycle = ", cpu.cycle)
	fmt.Println("memAddrReg = ", cpu.memAddrReg)
	fmt.Printf("ins reg = %08b\n", cpu.instructionReg)
	fmt.Println("REG_A = ", cpu.regA)
	fmt.Println("REG_B = ", cpu.regB)

	// Loading instruction into the instruction register
	var ins byte = (cpu.instructionReg >> 4) & 0xF
	fmt.Printf("ins = %04b\n", ins)
	var moff byte = ins*INSTRUCTION_LENGTH + cpu.cycle
	cpu.runMicrocode(microcode[moff])

	// Increment op counter
	cpu.cycle = (cpu.cycle + 1) % 8

	if cpu.cycle == 0 {
		fmt.Println("========================================")
	}
}

func (cpu *CPU) runMicrocode(op int64) {
	if (op & C_CI) > 0 {
		color.Red("C_CI")
		cpu.pcInc()
	}
	if (op & C_CO) > 0 {
		color.Red("C_CO")
		cpu.pcOut()
	}
	if (op & C_RO) > 0 {
		color.Red("C_RO")
		cpu.memOut()
	}
	if (op & C_IO) > 0 {
		color.Red("C_IO")
		cpu.insRegOut()
	}
	if (op & C_MI) > 0 {
		color.Red("C_MI")
		cpu.memAddrIn()
	}
	if (op & C_II) > 0 {
		color.Red("C_II")
		cpu.insRegIn()
	}
	if (op & C_EO) > 0 {
		color.Red("C_EO")
		cpu.aluOut()
	}
	if (op & C_OI) > 0 {
		color.Red("OI")
		cpu.out()
	}
	if (op & C_AI) > 0 {
		color.Red("AI")
		cpu.regaIn()
	}
	if (op & C_AO) > 0 {
		color.Red("AO")
		cpu.regaOut()
	}
	if (op & C_BI) > 0 {
		color.Red("BI")
		cpu.regbIn()
	}
	if (op & C_BO) > 0 {
		color.Red("BO")
		cpu.regbOut()
	}
	if (op & C_HALT) > 0 {
		color.Red("C_HALT")
		cpu.disableClock()
	}
}

func (cpu *CPU) disableClock() {
	cpu.clockEnabled = false
}

func (cpu *CPU) out() {
	color.Cyan("OUT : %08b (%d)", cpu.bus, cpu.bus)
}

func (cpu *CPU) aluOut() {
	cpu.bus = (cpu.regA + cpu.regB) & 0xFF
}

func (cpu *CPU) pcInc() {
	cpu.pc++
}

func (cpu *CPU) pcOut() {
	cpu.bus = cpu.pc
}

func (cpu *CPU) memAddrIn() {
	cpu.memAddrReg = cpu.bus & 0x0F
}

func (cpu *CPU) memOut() {
	fmt.Println("DEBUG - mem addr : ", cpu.memAddrReg)
	cpu.memory.Dump()
	cpu.bus = cpu.memory.Read(cpu.memAddrReg)
}

func (cpu *CPU) insRegIn() {
	cpu.instructionReg = cpu.bus
}

func (cpu *CPU) insRegOut() {
	cpu.bus = cpu.instructionReg & 0x0F
}

func (cpu *CPU) regaIn() {
	cpu.regA = cpu.bus
}

func (cpu *CPU) regbIn() {
	cpu.regB = cpu.bus
}

func (cpu *CPU) regbOut() {
	cpu.regB = cpu.bus
}

func (cpu *CPU) regaOut() {
	cpu.bus = cpu.regA
}
