package models

type ShortLinkRedisData struct {
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}

func (d *ShortLinkRedisData) GetLongURL() string  { return d.LongURL }
func (d *ShortLinkRedisData) GetShortURL() string { return d.ShortURL }
