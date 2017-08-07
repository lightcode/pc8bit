package main

import "fmt"
import "time"
import "github.com/fatih/color"

const INSTRUCTION_LENGTH = 8

const (
	NOP    = iota
	PC_INC = 1 << iota
	PC_OUT
	MEM_ADDR_IN
	MEM_OUT
	INS_REG_IN
	INS_REG_OUT
	REGA_IN
	REGB_IN
	ALU_OUT
	OP_OUT
	REGA_OUT
	CLOCK_HALT
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
	if cpu.cycle == 0 {
		cpu.runMicrocode(PC_OUT | MEM_ADDR_IN)
	} else if cpu.cycle == 1 {
		cpu.runMicrocode(PC_INC | MEM_OUT | INS_REG_IN)
	} else {
		var ins byte = (cpu.instructionReg >> 4) & 0xF
		fmt.Printf("ins = %04b\n", ins)
		var moff byte = ins*INSTRUCTION_LENGTH + cpu.cycle
		cpu.runMicrocode(microcode[moff])
	}

	// Increment op counter
	cpu.cycle = (cpu.cycle + 1) % 8

	if cpu.cycle == 0 {
		fmt.Println("========================================")
	}
}

func (cpu *CPU) runMicrocode(op int64) {
	if (op & NOP) > 0 {
		fmt.Println("NOP")
		cpu.nop()
	}
	if (op & PC_INC) > 0 {
		color.Red("PC_INC")
		cpu.pcInc()
	}
	if (op & PC_OUT) > 0 {
		color.Red("PC_OUT")
		cpu.pcOut()
	}
	if (op & MEM_OUT) > 0 {
		color.Red("MEM_OUT")
		cpu.memOut()
	}
	if (op & INS_REG_OUT) > 0 {
		color.Red("INS_REG_OUT")
		cpu.insRegOut()
	}
	if (op & MEM_ADDR_IN) > 0 {
		color.Red("MEM_ADDR_IN")
		cpu.memAddrIn()
	}
	if (op & INS_REG_IN) > 0 {
		color.Red("INS_REG_IN")
		cpu.insRegIn()
	}
	if (op & ALU_OUT) > 0 {
		color.Red("ALU_OUT")
		cpu.aluOut()
	}
	if (op & OP_OUT) > 0 {
		color.Red("OUT")
		cpu.out()
	}
	if (op & REGA_IN) > 0 {
		color.Red("REGA_IN")
		cpu.regaIn()
	}
	if (op & REGA_OUT) > 0 {
		color.Red("REGA_OUT")
		cpu.regaOut()
	}
	if (op & REGB_IN) > 0 {
		color.Red("REGB_IN")
		cpu.regbIn()
	}
	if (op & CLOCK_HALT) > 0 {
		color.Red("CLOCK_HALT")
		cpu.clockHalt()
	}
}

func (cpu *CPU) nop() {

}

func (cpu *CPU) clockHalt() {
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

func (cpu *CPU) regaOut() {
	cpu.bus = cpu.regA
}
