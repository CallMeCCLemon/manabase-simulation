package main

type Card struct {
	Land    *Land
	NonLand *NonLand
}

func NewCard(land *Land, nonLand *NonLand) *Card {
	return &Card{
		Land:    land,
		NonLand: nonLand,
	}
}
