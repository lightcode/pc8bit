package main

const (
	I_NOP = iota & 0xF
	I_LDA
	I_ADD
	I_OUT
	I_JMP
	I_HLT
)

const (
	CY_1 = C_CO | C_MI
	CY_2 = C_CI | C_RO | C_II
	CY_3 = C_CO | C_MI
	CY_4 = C_CI | C_RO | C_ZI
)

var microcode = []Opcode{
	CY_1, CY_2, CY_3, CY_4, 0, 0, 0, 0, // NOP
	CY_1, CY_2, CY_3, CY_4, C_ZO | C_MI, C_RO | C_AI, 0, 0, // LDA
	CY_1, CY_2, CY_3, CY_4, C_ZO | C_MI, C_RO | C_BI, C_EO | C_AI, 0, // ADD
	CY_1, CY_2, CY_3, CY_4, C_AO, C_OI, 0, 0, // OUT
	CY_1, CY_2, CY_3, CY_4, C_ZO | C_J, 0, 0, 0, // JMP
	CY_1, CY_2, CY_3, CY_4, C_HALT, 0, 0, 0, // HLT
}
