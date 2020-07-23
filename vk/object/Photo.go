package object

type PhotoSize struct {
	Type   string `json:"type" map:"type"`
	Url    string `json:"url" map:"url"`
	Width  int    `json:"width" map:"width"`
	Height int    `json:"height" map:"height"`
}

type Photo struct {
	ID      int         `json:"id" map:"id"`
	AlbumID int         `json:"album_id" map:"album_id"`
	OwnerID int         `json:"owner_id" map:"owner_id"`
	Text    string      `json:"text" map:"text"`
	Date    int         `json:"date" map:"date"`
	Size    []PhotoSize `json:"sizes" map:"sizes"`
	Width   int         `json:"width" map:"width"`
	Height  int         `json:"height" map:"height"`

	biggerImageUrl string
}

func (p *Photo) GetBiggerImageUrl() (currentUrl string) {
	if p.biggerImageUrl != "" {
		return p.biggerImageUrl
	}

	var maxSize int = 0
	for _, size := range p.Size {
		crntSize := size.Width * size.Height
		if crntSize > maxSize {
			currentUrl = size.Url
			maxSize = crntSize
		}
	}

	p.biggerImageUrl = currentUrl
	return
}
