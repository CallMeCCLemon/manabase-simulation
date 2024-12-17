package validation

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/pgtype"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"manabase-simulation/package/model"
	"time"
)

type Land struct {
	gorm.Model
	Name              string       `gorm:"uniqueIndex"`
	Colors            pgtype.JSONB `gorm:"type:jsonb"`
	EntersTapped      bool
	Types             pgtype.JSONB `gorm:"type:jsonb"`
	UntappedCondition pgtype.JSONB `gorm:"type:jsonb"`
	ActivationCost    pgtype.JSONB `gorm:"type:jsonb"`
}

func toGormModel(card *model.Card) (*Land, error) {
	l := &Land{
		Name:         card.Name,
		EntersTapped: card.Land.EntersTapped,
	}
	err := l.Colors.Set(card.Land.Colors)
	if err != nil {
		return nil, err
	}
	err = l.Types.Set(card.Land.Types)
	if err != nil {
		return nil, err
	}
	err = l.UntappedCondition.Set(card.Land.UntappedCondition)
	if err != nil {
		return nil, err
	}
	err = l.ActivationCost.Set(card.Land.ActivationCost)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Land) Get() (*model.Land, error) {
	var untappedCondition *model.UntappedCondition
	var activationCost *model.ActivationCost
	var colors []model.ManaColor
	var landTypes []model.LandType

	err := l.UntappedCondition.AssignTo(&untappedCondition)
	if err != nil {
		return nil, err
	}

	err = l.ActivationCost.AssignTo(&activationCost)
	if err != nil {
		return nil, err
	}
	err = l.Colors.AssignTo(&colors)
	if err != nil {
		return nil, err
	}

	err = l.Types.AssignTo(&landTypes)
	if err != nil {
		return nil, err
	}

	return &model.Land{
		Name:              l.Name,
		Colors:            colors,
		EntersTapped:      l.EntersTapped,
		Types:             landTypes,
		UntappedCondition: untappedCondition,
		ActivationCost:    activationCost,
	}, nil
}

type CardDbAccessor interface {
	// CreateTables creates all necessary tables for the card database to work.
	CreateTables() error

	// GetCard returns a card by name.
	GetCard(name string) (*model.Card, error)

	// GetCards returns a list of cards. This is more efficient for bulk lookups of cards.
	GetCards(names []string) ([]model.Card, error)

	// WriteCard writes a card.
	WriteCard(card *model.Card) (int64, error)

	// WriteCards writes a list of cards.
	WriteCards(cards []model.Card) (int64, error)
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
	err := c.GormDB.AutoMigrate(&Land{})
	return err
}

func (c *CardDbAccessorImpl) GetCard(name string) (*model.Card, error) {
	var card *Land
	res := c.GormDB.Where("name = ?", name).First(&card)
	if res.Error != nil {
		return nil, res.Error
	}

	var colors []string
	_ = json.Unmarshal(card.Colors.Bytes, &colors)
	println(fmt.Sprintf("Colors: %s", string(card.Colors.Bytes)))
	err := card.Colors.AssignTo(&colors)
	if err != nil {
		return nil, err
	}

	var untappedCondition *model.UntappedCondition
	err = card.UntappedCondition.Set(untappedCondition)
	if err != nil {
		return nil, err
	}

	var activationCost *model.ActivationCost
	err = card.ActivationCost.Set(activationCost)
	if err != nil {
		return nil, err
	}

	var types []model.LandType
	err = card.Types.Set(&types)
	if err != nil {
		return nil, err
	}

	l := &model.Land{
		Name: card.Name,
		//Colors:            colors,
		EntersTapped:      card.EntersTapped,
		Types:             types,
		UntappedCondition: untappedCondition,
		ActivationCost:    activationCost,
	}

	return &model.Card{
		Name:    card.Name,
		Land:    l,
		NonLand: nil,
	}, nil
}

func (c *CardDbAccessorImpl) GetCards(types []string) ([]model.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CardDbAccessorImpl) WriteCard(card *model.Card) (int64, error) {
	gormCard, err := toGormModel(card)
	if err != nil {
		return 0, err
	}
	result := c.GormDB.Save(gormCard)
	return result.RowsAffected, result.Error
}

func (c *CardDbAccessorImpl) WriteCards(cards []model.Card) (int64, error) {
	gormCards := make([]*Land, len(cards))
	for i, card := range cards {
		gormCard, err := toGormModel(&card)
		if err != nil {
			return 0, err
		}

		gormCards[i] = gormCard
	}

	res := c.GormDB.Save(gormCards)

	return res.RowsAffected, res.Error
}
