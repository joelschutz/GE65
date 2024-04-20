package main

import (
	"fmt"
)

type Address uint16
type Word uint8
type ComponentType uint8

const (
	RO ComponentType = iota // If a component is read-only
	WO                      // If a component is write-only
	RW                      // If a component is read-write
)

type Memory struct {
	rMemory map[Address]Reader
	wMemory map[Address]Writer
}

func (m Memory) Read(addr Address) (Word, error) {
	return m.rMemory[addr].Read(addr)
}

func (m Memory) Write(addr Address, data Word) error {
	return m.wMemory[addr].Write(addr, data)
}

func (m *Memory) AddReader(addr Address, size uint16, c Component) error {
	r, ok := c.(Reader)
	if ok {
		for i := addr; i < addr+Address(size); i++ {
			m.rMemory[i] = r
		}
	}
	return fmt.Errorf("Component is not a reader")
}

func (m *Memory) AddWriter(addr Address, size uint16, c Component) error {
	w, ok := c.(Writer)
	if ok {
		for i := addr; i < addr+Address(size); i++ {
			m.wMemory[i] = w
		}
	}
	return fmt.Errorf("Component is not a writer")
}

func (m *Memory) PlugComponent(t ComponentType, addr Address, size uint16, c Component) error {
	if c.Init(addr) != nil {
		return fmt.Errorf("Error initializing component")
	}

	switch t {
	case RO:
		return m.AddReader(addr, size, c)
	case WO:
		return m.AddWriter(addr, size, c)
	case RW:
		err1 := m.AddReader(addr, size, c)
		err2 := m.AddWriter(addr, size, c)

		if err1 != nil || err2 != nil {
			return fmt.Errorf(err1.Error() + ":" + err2.Error())
		}
	}

	return nil
}

// beevik/go6502 interface

// LoadByte loads a single byte from the address and returns it.
func (m Memory) LoadByte(addr uint16) byte {
	return byte(m.LoadAddress(addr))
}

// LoadBytes loads multiple bytes from the address and stores them into
// the buffer 'b'.
func (m Memory) LoadBytes(addr uint16, b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = m.LoadByte(addr + uint16(i))
	}
}

// LoadAddress loads a 16-bit address value from the requested address and
// returns it.
func (m Memory) LoadAddress(addr uint16) uint16 {
	d, err := m.Read(Address(addr))
	if err != nil {
		panic(err)
	}
	return uint16(d)
}

// StoreByte stores a byte to the requested address.
func (m Memory) StoreByte(addr uint16, v byte) {
	m.StoreAddress(addr, uint16(v))
}

// StoreBytes stores multiple bytes to the requested address.
func (m Memory) StoreBytes(addr uint16, b []byte) {
	for i := 0; i < len(b); i++ {
		m.StoreByte(addr+uint16(i), b[i])
	}
}

// StoreAddres stores a 16-bit address 'v' to the requested address.
func (m Memory) StoreAddress(addr uint16, v uint16) {
	err := m.Write(Address(addr), Word(v))
	if err != nil {
		panic(err)
	}
}
