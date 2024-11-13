package model

import (
	"errors"
	"manabase-simulation/package/reader"
)

// CardDB Holds all the known MTG cards from Scryfall.
type CardDB struct {
	data map[string]reader.ScryfallCard
}

// NewCardDB Creates a new CardDB instance.
func NewCardDB() (*CardDB, error) {
	data, err := reader.ReadScryfallDataJSONFile("../../data/scryfall-db.json")
	if err != nil {
		return nil, err
	}
	return &CardDB{
		data: data,
	}, nil
}

func (s *CardDB) GetCard(name string) (*reader.ScryfallCard, error) {
	card, ok := s.data[name]
	if !ok {
		return nil, errors.New("card not found")
	}
	return &card, nil
}
