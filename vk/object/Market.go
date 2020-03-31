package object

type Price struct {
	Amount string `json:"amount" map:"amount"`
	Text   string `json:"text" map:"text"`
}

type Market struct {
	ID          float64 `json:"id" map:"id"`
	OwnerID     float64 `json:"owner_id" map:"owner_id"`
	Title       string  `json:"title" map:"title"`
	Description string  `json:"description" map:"description"`
	Price       *Price  `json:"price" map:"price"`
	Photo       string  `json:"thumb_photo" map:"thumb_photo"`
	price       float64
}

//func (p *Market) GetPrice() (price float64) {
//	if p.price != 0 {
//		return p.price
//	}
//	p.price = pyraconv.ToFloat64(p.Price.Amount)
//	price = p.price
//	return
//}
