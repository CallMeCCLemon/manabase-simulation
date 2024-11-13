package reader

import (
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
	cards, err := ReadJSONFile[[]ScryfallCard](filename)
	if err != nil {
		return nil, err
	}

	cardsLookup := make(map[string]ScryfallCard, len(cards))
	for _, card := range cards {
		cardsLookup[card.Name] = card
	}
	return cardsLookup, nil
}
