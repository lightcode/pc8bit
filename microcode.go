package main

const (
	I_NOP = iota & 0xF
	I_LDA
	I_ADD
	I_OUT
	I_HLT
)

var microcode = []int64{
	C_CO | C_MI, C_CI | C_RO | C_II, 0, 0, 0, 0, 0, 0, // NOP
	C_CO | C_MI, C_CI | C_RO | C_II, C_IO | C_MI, C_RO | C_AI, 0, 0, 0, 0, // LDA
	C_CO | C_MI, C_CI | C_RO | C_II, C_IO | C_MI, C_RO | C_BI, C_EO | C_AI, 0, 0, 0, // ADD
	C_CO | C_MI, C_CI | C_RO | C_II, C_AO, C_OI, 0, 0, 0, 0, // OUT
	C_CO | C_MI, C_CI | C_RO | C_II, C_HALT, 0, 0, 0, 0, 0, // HLT
}
