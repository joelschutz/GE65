package main

import "github.com/beevik/go6502/cpu"

func main() {
	// Create the memory
	m := Memory{
		rMemory: map[Address]Reader{},
		wMemory: map[Address]Writer{},
	}

	// Create the RAM
	ram := NewRAM(256)

	m.PlugComponent(RW, Address(0), 256, ram)

	// Create the CPU
	c65 := cpu.NewCPU(cpu.CMOS, m)
	c65.irq
}
