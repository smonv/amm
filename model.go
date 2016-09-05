package amm

// AlternativeTitle ...
type AlternativeTitle struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
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
	Title             string             `json:"title"`
	AlternativeTitles []AlternativeTitle `json:"alternative_titles"`
	Type              string             `json:"type"`
	Episodes          int                `json:"episodes"`
	Subtitle          []Subtitle         `json:"subtitles"`
	Extra             []string           `json:"extra"`
}
