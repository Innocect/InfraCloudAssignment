package models

type ShortLinkMongoData struct {
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}

func (d *ShortLinkMongoData) GetLongURL() string  { return d.LongURL }
func (d *ShortLinkMongoData) GetShortURL() string { return d.ShortURL }
