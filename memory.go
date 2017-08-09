package main

import "fmt"

type Memory struct {
	data [256]byte
}

func (m *Memory) Read(addr byte) byte {
	return m.data[int(addr)]
}

func (m *Memory) Write(addr, data byte) {
	m.data[int(addr)] = data
}

func (m *Memory) Dump() {
	fmt.Println(m.data)
}
