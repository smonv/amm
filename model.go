package amm

// Title ...
type Title struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

// NewTitle ...
func NewTitle(args []string) *Title {
	var t string
	var l string

	switch len(args) {
	case 1:
		t = args[0]
	case 2:
		t = args[0]
		l = args[1]
	default:
		return nil
	}

	return &Title{
		Text: t,
		Lang: l,
	}
}

// Subtitle ...
type Subtitle struct {
	Lang   string   `json:"lang"`
	Fansub string   `json:"fansub"`
	Encode []string `json:"encode"`
	Source []string `json:"source"`
}

// Anime ...
type Anime struct {
	Title             Title      `json:"title"`
	AlternativeTitles []Title    `json:"alternative_titles"`
	Synonyms          []string   `json:"synonyms"`
	Type              string     `json:"type"`
	Episodes          int        `json:"episodes"`
	Subtitle          []Subtitle `json:"subtitles"`
	Extra             []string   `json:"extra"`
}
