package validation

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"manabase-simulation/package/model"
	"time"
)

type GormCard struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex"`
	Land    string
	NonLand string
}

func toGormCard(card *model.Card) (*GormCard, error) {
	var land []byte
	land, err := json.Marshal(card.Land)
	if err != nil {
		return nil, err
	}

	var nonLand []byte
	nonLand, err = json.Marshal(card.NonLand)
	if err != nil {
		return nil, err
	}

	return &GormCard{
		Name:    card.Name,
		Land:    string(land),
		NonLand: string(nonLand),
	}, nil
}

func toModelCard(card *GormCard) (*model.Card, error) {
	var land *model.Land
	err := json.Unmarshal([]byte(card.Land), &land)
	if err != nil {
		return nil, err
	}

	var nonLand *model.NonLand
	err = json.Unmarshal([]byte(card.NonLand), &nonLand)
	if err != nil {
		return nil, err
	}

	return &model.Card{
		Name:    card.Name,
		Land:    land,
		NonLand: nonLand,
	}, nil
}

// CardDbAccessor is an interface for accessing the card database.
type CardDbAccessor interface {
	// CreateTables creates all necessary tables for the card database to work.
	CreateTables() error

	// GetCard returns a card from the database.
	GetCard(name string) (*model.Card, error)

	// WriteCard writes a single card to the database.
	WriteCard(card *model.Card) (int64, error)

	// WriteCards writes a batch of cards to the database.
	WriteCards(card []*model.Card) (int64, error)
}

var _ CardDbAccessor = &CardDbAccessorImpl{}

type CardDbAccessorImpl struct {
	GormDB *gorm.DB
}

// NewPsqlDbAccessor Creates a new CardDbAccessor instance with a psql connection.
func NewPsqlDbAccessor(config postgres.Config) (*CardDbAccessorImpl, error) {
	psql := postgres.New(config)
	return NewCardDbAccessor(psql)
}

// NewCardDbAccessor Creates a new CardDbAccessor instance.
func NewCardDbAccessor(d gorm.Dialector) (*CardDbAccessorImpl, error) {
	gormCfg := gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	gormDB, err := gorm.Open(d, &gormCfg)
	if err != nil {
		return nil, err
	}

	return &CardDbAccessorImpl{
		GormDB: gormDB,
	}, nil
}

func (c *CardDbAccessorImpl) CreateTables() error {
	return c.GormDB.AutoMigrate(&GormCard{})
}

func (c *CardDbAccessorImpl) GetCard(name string) (*model.Card, error) {
	var card *GormCard
	res := c.GormDB.Where("name = ?", name).First(&card)
	if res.Error != nil {
		return nil, res.Error
	}

	return toModelCard(card)
}

func (c *CardDbAccessorImpl) WriteCard(card *model.Card) (int64, error) {
	gormCard, err := toGormCard(card)
	if err != nil {
		return 0, err
	}
	result := c.GormDB.Save(gormCard)
	return result.RowsAffected, result.Error
}

func (c *CardDbAccessorImpl) WriteCards(cards []*model.Card) (int64, error) {
	gormCards := make([]*GormCard, len(cards))
	for i, card := range cards {
		gormCard, err := toGormCard(card)
		if err != nil {
			return 0, err
		}

		gormCards[i] = gormCard
	}

	res := c.GormDB.Save(gormCards)

	return res.RowsAffected, res.Error
}
