package main

type RWer interface {
	Reader
	Writer
}

type Reader interface {
	Read(addr Address) (Word, error)
}

type Writer interface {
	Write(addr Address, data Word) error
}

type Component interface {
	Init(location Address) error
}
