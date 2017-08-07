package main

type Opcode int

var mins = []Opcode{
	C_HALT,
	C_CO, C_RO, C_IO, C_AO, C_BO, C_EO, // Always put output before input
	C_CI, C_MI, C_II, C_OI, C_AI, C_BI,
}

// Micro instructions
const (
	_      Opcode = iota
	C_CI          = 1 << iota // program Counter Increment
	C_CO                      // program Counter Out
	C_RO                      // RAM Out
	C_II                      // Instruction register In
	C_IO                      // Instruction register Out
	C_EO                      // ALU Out
	C_AI                      // Register A In
	C_AO                      // Register A Out
	C_BI                      // Register B In
	C_BO                      // Register B Out
	C_HALT                    // Halt: disable the clock
	C_MI                      // Memory Register In
	C_OI                      // Out In: bus -> out
)
