package main

type Handler interface {
	Add() error
	Del() error
	Mod() error
	List() error
}

type Mux struct {
}

func NewMux() *Mux {
	return &Mux{}
}

func (m *Mux) Register(pattern string, handler Handler) {

}
