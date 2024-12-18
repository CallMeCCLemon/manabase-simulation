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

type GormLand struct {
	gorm.Model
	Name              string `gorm:"uniqueIndex"`
	EntersTapped      bool
	Colors            string
	Types             string
	UntappedCondition string
	ActivationCost    string
}

type GormNonLand struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	CastingCost string
	Quantity    int
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

func toGormLand(land *model.Land) (*GormLand, error) {
	c, err := json.Marshal(land.Colors)
	if err != nil {
		return nil, err
	}

	t, err := json.Marshal(land.Types)
	if err != nil {
		return nil, err
	}

	u, err := json.Marshal(land.UntappedCondition)
	if err != nil {
		return nil, err
	}

	ac, err := json.Marshal(land.ActivationCost)
	if err != nil {
		return nil, err
	}
	l := &GormLand{
		Name:              land.Name,
		EntersTapped:      land.EntersTapped,
		Colors:            string(c),
		Types:             string(t),
		UntappedCondition: string(u),
		ActivationCost:    string(ac),
	}

	return l, nil
}

func toGormNonLand(nonLand *model.NonLand) (*GormNonLand, error) {
	c, err := json.Marshal(nonLand.CastingCost)
	if err != nil {
		return nil, err
	}

	return &GormNonLand{
		Name:        nonLand.Name,
		CastingCost: string(c),
		Quantity:    nonLand.Quantity,
	}, nil
}

type CardDbAccessor interface {
	// CreateTables creates all necessary tables for the card database to work.
	CreateTables() error

	GetCard(name string) (*model.Card, error)

	WriteCard(card *model.Card) (int64, error)
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
	err := c.GormDB.AutoMigrate(&GormLand{})
	if err != nil {
		return err
	}

	err = c.GormDB.AutoMigrate(&GormCard{})
	if err != nil {
		return err
	}

	err = c.GormDB.AutoMigrate(&GormNonLand{})
	return err
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

func (c *CardDbAccessorImpl) WriteLands(lands []model.Land) (int64, error) {
	gormLands := make([]*GormLand, len(lands))
	for i, card := range lands {
		gormCard, err := toGormLand(&card)
		if err != nil {
			return 0, err
		}

		gormLands[i] = gormCard
	}

	res := c.GormDB.Save(gormLands)

	return res.RowsAffected, res.Error
}

func (c *CardDbAccessorImpl) WriteNonLands(nonLands []model.NonLand) (int64, error) {
	gormNonLands := make([]*GormNonLand, len(nonLands))
	for i, card := range nonLands {
		gormNonLand, err := toGormNonLand(&card)
		if err != nil {
			return 0, err
		}

		gormNonLands[i] = gormNonLand
	}

	res := c.GormDB.Save(gormNonLands)

	return res.RowsAffected, res.Error
}
