package model

type DeckStats struct {
	TotalCards    int32    `json:"totalCards"`
	Lands         int32    `json:"lands"`
	NonLands      int32    `json:"nonLands"`
	TotalManaPips ManaCost `json:"totalManaPips"`
}
