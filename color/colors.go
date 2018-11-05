package main

import (
	"sync"
)

//Color type
type Color struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type ColorStore struct {
	sync.Mutex
	Colors []Color
}

func (cs *ColorStore) GetAll() []Color {
	cs.Lock()
	defer cs.Unlock()
	return cs.Colors
}

func (cs *ColorStore) Set(color Color) {
	cs.Lock()
	defer cs.Unlock()
	cs.Colors = append(cs.Colors, color)
}

func NewColors() *ColorStore {
	var cs ColorStore
	cs.Colors = NewThreeColors()
	return &cs
}

func NewThreeColors() []Color {
	return []Color{
		Color{
			Name: "red",
			Hex:  "#FF0000",
		},
		Color{
			Name: "green",
			Hex:  "#00FF00",
		},
		Color{
			Name: "blue",
			Hex:  "#0000FF",
		},
	}
}
