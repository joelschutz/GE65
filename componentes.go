package main

import "fmt"

type ROM struct {
	zLocation Address // the zero-based location of the ROM
	data      []byte
}

func (r *ROM) Init(location Address) error {
	// Check if data is present
	if r.data == nil {
		return fmt.Errorf("No data present in ROM")
	}

	r.zLocation = location
	return nil
}

func (r *ROM) Read(addr Address) (Word, error) {
	// Calculate the local address
	localAddr := addr - r.zLocation

	// Get the data
	d := r.data[localAddr]

	return Word(d), nil
}

func NewROM(data []byte) *ROM {
	return &ROM{
		data: data,
	}
}

type RAM struct {
	zLocation Address // the zero-based location of the RAM
	data      []byte
}

func (r *RAM) Init(location Address) error {
	// Check if data is present
	if r.data == nil {
		return fmt.Errorf("No space present in RAM")
	}

	// Set the location
	r.zLocation = location
	return nil
}

func (r *RAM) Read(addr Address) (Word, error) {
	// Calculate the local address
	localAddr := addr - r.zLocation

	// Get the data
	d := r.data[localAddr] // Probably should recover from this

	return Word(d), nil
}

func (r *RAM) Write(addr Address, data Word) error {
	// Calculate the local address
	localAddr := addr - r.zLocation

	r.data[localAddr] = byte(data) // Probably should recover from this
	return nil
}

func NewRAM(size uint16) *RAM {
	return &RAM{
		data: make([]byte, size),
	}
}
