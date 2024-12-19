package scryfall

import (
	"fmt"
	"manabase-simulation/package/model"
	"manabase-simulation/package/reader"
	"manabase-simulation/package/validation"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ScryfallCard Represents a card from Scryfall.
type ScryfallCard struct {
	Object          string            `json:"object"`
	ID              string            `json:"id"`
	OracleID        string            `json:"oracle_id"`
	MultiverseIDs   []int             `json:"multiverse_ids"`
	MtgoID          int               `json:"mtgo_id"`
	MtgoFoilID      int               `json:"mtgo_foil_id"`
	TcgplayerID     int               `json:"tcgplayer_id"`
	CardmarketID    int               `json:"cardmarket_id"`
	Name            string            `json:"name"`
	Lang            string            `json:"lang"`
	ReleasedAt      CustomDate        `json:"released_at"`
	URI             string            `json:"uri"`
	ScryfallURI     string            `json:"scryfall_uri"`
	Layout          string            `json:"layout"`
	HighresImage    bool              `json:"highres_image"`
	ImageStatus     string            `json:"image_status"`
	ImageUris       ImageURIs         `json:"image_uris"`
	ManaCost        string            `json:"mana_cost"`
	Cmc             float64           `json:"cmc"`
	TypeLine        string            `json:"type_line"`
	OracleText      string            `json:"oracle_text"`
	Colors          []string          `json:"colors"`
	ColorIdentity   []string          `json:"color_identity"`
	Keywords        []string          `json:"keywords"`
	Legalities      map[string]string `json:"legalities"`
	Games           []string          `json:"games"`
	Reserved        bool              `json:"reserved"`
	Foil            bool              `json:"foil"`
	Nonfoil         bool              `json:"nonfoil"`
	Finishes        []string          `json:"finishes"`
	Oversized       bool              `json:"oversized"`
	Promo           bool              `json:"promo"`
	Reprint         bool              `json:"reprint"`
	Variation       bool              `json:"variation"`
	SetID           string            `json:"set_id"`
	Set             string            `json:"set"`
	SetName         string            `json:"set_name"`
	SetType         string            `json:"set_type"`
	SetURI          string            `json:"set_uri"`
	SetSearchURI    string            `json:"set_search_uri"`
	ScryfallSetURI  string            `json:"scryfall_set_uri"`
	RulingsURI      string            `json:"rulings_uri"`
	PrintsSearchURI string            `json:"prints_search_uri"`
	CollectorNumber string            `json:"collector_number"`
	Digital         bool              `json:"digital"`
	Rarity          string            `json:"rarity"`
	FlavorText      string            `json:"flavor_text"`
	CardBackID      string            `json:"card_back_id"`
	Artist          string            `json:"artist"`
	ArtistIDs       []string          `json:"artist_ids"`
	IllustrationID  string            `json:"illustration_id"`
	BorderColor     string            `json:"border_color"`
	Frame           string            `json:"frame"`
	FullArt         bool              `json:"full_art"`
	Textless        bool              `json:"textless"`
	Booster         bool              `json:"booster"`
	StorySpotlight  bool              `json:"story_spotlight"`
	EdhrecRank      int               `json:"edhrec_rank"`
	Prices          Prices            `json:"prices"`
	RelatedUris     map[string]string `json:"related_uris"`
	PurchaseUris    map[string]string `json:"purchase_uris"`
	ProducedMana    []string          `json:"produced_mana,omitempty"`
}

type ImageURIs struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type Prices struct {
	USD       *string `json:"usd"`
	USDFoil   *string `json:"usd_foil"`
	USDEtched *string `json:"usd_etched"`
	EUR       *string `json:"eur"`
	EURFoil   *string `json:"eur_foil"`
	Tix       *string `json:"tix"`
}

type CustomDate struct {
	time.Time
}

const customDateFormat = "2006-01-02"

func (c *CustomDate) UnmarshalJSON(b []byte) error {
	// Trim the quotes from the JSON string
	str := string(b)
	str = str[1 : len(str)-1]

	// Parse the date using the custom layout
	t, err := time.Parse(customDateFormat, str)
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

// ReadScryfallDataJSONFile Reads a Scryfall JSON Dump file.
func ReadScryfallDataJSONFile(filename string) (map[string]ScryfallCard, error) {
	var cards []ScryfallCard
	cards, err := reader.ReadJSONFile[[]ScryfallCard](filename)
	if err != nil {
		return nil, err
	}

	cardsLookup := make(map[string]ScryfallCard, len(cards))
	for _, card := range cards {
		cardsLookup[card.Name] = card
	}
	return cardsLookup, nil
}

// WriteNonLandsToDB Reads all the non-lands from Scryfall and uploads them to the database in a parsed format.
func WriteNonLandsToDB(accessor validation.CardDbAccessor, nonLands []ScryfallCard, writeToDB bool) ([]*model.Card, error) {
	parsedNonLands := make([]*model.Card, len(nonLands))
	for i, nonLand := range nonLands {
		card := parseNonLandCard(nonLand)
		if writeToDB {
			_, err := accessor.WriteCard(parseNonLandCard(nonLand))
			if err != nil {
				return parsedNonLands, err
			}
		}
		println(fmt.Sprintf("Adding %s", card.Name))
		parsedNonLands[i] = card
	}
	return parsedNonLands, nil
}

func parseNonLandCard(card ScryfallCard) *model.Card {
	nonLand := &model.NonLand{
		Name:        card.Name,
		CastingCost: parseNonLandCastingCost(card.ManaCost),
		Quantity:    1,
	}

	return &model.Card{
		Name:    card.Name,
		Land:    nil,
		NonLand: nonLand,
	}

}

func parseNonLandCastingCost(cost string) model.ManaCost {
	re := regexp.MustCompile(`\{([^{}]*)\}`)
	matches := re.FindAllStringSubmatch(cost, -1)

	manaCost := model.ManaCost{}

	// Extract and print the captured groups
	for _, m := range matches {
		if len(m) > 1 {
			// Kind of a funky way to determine if it's a number, but whatever.
			if i, err := strconv.Atoi(m[1]); err == nil {
				manaCost.GenericCost = i
			} else {
				color := charToManaColor(m[1])
				manaCost.ColorRequirements = append(manaCost.ColorRequirements, color)
			}
		}
	}

	return manaCost
}

// WriteLandsToDB Reads all the lands from Scryfall and uploads them to the database in a parsed format.
func WriteLandsToDB(accessor validation.CardDbAccessor, lands []ScryfallCard, writeToDB bool) ([]*model.Card, error) {
	parsedLands := make([]*model.Card, len(lands))
	for i, land := range lands {
		card := parseLandCard(land)
		if writeToDB {
			_, err := accessor.WriteCard(parseLandCard(land))
			if err != nil {
				return parsedLands, err
			}
		}
		println(fmt.Sprintf("Adding %s", card.Name))
		parsedLands[i] = card
	}
	return parsedLands, nil
}

// parseLandCard parses a land card from Scryfall and returns a model.Card.
func parseLandCard(card ScryfallCard) *model.Card {
	l := model.Land{
		Name:         card.Name,
		Colors:       make([]model.ManaColor, 0),
		EntersTapped: false,
		Quantity:     1,
		Types:        make([]model.LandType, 0),
	}
	for _, manaColor := range card.ProducedMana {
		l.Colors = append(l.Colors, charToManaColor(manaColor))
	}
	l.Types = parseLandTypes(card.TypeLine)

	entersTapped, untappedCondition := parseOracleTextForEntersTapped(l.Name, card.OracleText)
	l.EntersTapped = entersTapped
	l.UntappedCondition = untappedCondition

	return &model.Card{
		Name:    l.Name,
		Land:    &l,
		NonLand: nil,
	}
}

func parseOracleTextForEntersTapped(name string, s string) (entersTapped bool, cond *model.UntappedCondition) {
	cond = nil
	entersTapped = false
	pattern := fmt.Sprintf("%s enters tapped", name)
	re := regexp.MustCompile(pattern)

	regexmatch := re.FindStringSubmatch(s)
	if len(regexmatch) >= 1 {
		entersTapped = true
		pattern = fmt.Sprintf("%s enters tapped unless (.*)", name)
		re = regexp.MustCompile(pattern)
		regexmatch = re.FindStringSubmatch(s)
		if len(regexmatch) > 1 {
			cond = &model.UntappedCondition{}
			switch regexmatch[1] {
			case "a player has 13 or less life.":
				cond.Type = model.UnluckyLand
			case "you control two or fewer other lands.":
				cond.Type = model.FastLand
			case "you revealed a Soldier card this way or you control a Soldier.":
				cond.Type = model.TypalLand
			case "you control a legendary green creature.":
				cond.Type = model.ArgothLand
			default:
				println(fmt.Sprintf("Unknown condition for %s", name))
				cond = nil
			}
		}
	}
	return entersTapped, cond
}

func charToManaColor(c string) model.ManaColor {
	switch c {
	case "W":
		return model.White
	case "U":
		return model.Blue
	case "B":
		return model.Black
	case "R":
		return model.Red
	case "G":
		return model.Green
	case "W/U":
		return model.Azorius
	case "W/B":
		return model.Orzhov
	case "W/R":
		return model.Boros
	case "W/G":
		return model.Selesnya
	case "U/B":
		return model.Dimir
	case "U/R":
		return model.Izzet
	case "U/G":
		return model.Simic
	case "B/R":
		return model.Rakdos
	case "B/G":
		return model.Golgari
	case "R/G":
		return model.Gruul
	default:
		return model.Colorless
	}
}

func parseLandTypes(s string) []model.LandType {
	parsedLandTypes := make([]model.LandType, 0)

	types := match(s)
	if types != nil {
		landTypes := strings.Split(strings.TrimSpace(*types), " ")
		for _, landType := range landTypes {
			lType := strToLandType(landType)
			if lType != nil {
				parsedLandTypes = append(parsedLandTypes, *lType)
			}
		}
	}
	return parsedLandTypes
}

func strToLandType(s string) *model.LandType {
	landType := model.LandType(s)
	return &landType
}

func match(s string) *string {
	pattern := `Land —\s*(.*)`
	re := regexp.MustCompile(pattern)

	regexmatch := re.FindStringSubmatch(s)
	if len(regexmatch) > 1 {
		return &regexmatch[1]
	}
	return nil
}
