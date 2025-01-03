package model

type DeckStats struct {
	TotalCards    int                `json:"totalCards"`
	Lands         int                `json:"lands"`
	NonLands      int                `json:"nonLands"`
	TotalManaPips SimplifiedManaCost `json:"totalManaPips"`
}
