package main

func NewInMemoryPlayerStore() *InMemoryPlayerScore {
	return &InMemoryPlayerScore{map[string]int{}}
}

type InMemoryPlayerScore struct {
	store map[string]int
}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerScore) RecordWin(name string) {
	i.store[name]++
}
