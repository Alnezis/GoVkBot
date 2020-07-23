package object

type Price struct {
	Amount string `json:"amount" map:"amount"`
	Text   string `json:"text" map:"text"`
}

type Market struct {
	ID          int    `json:"id" map:"id"`
	OwnerID     int    `json:"owner_id" map:"owner_id"`
	Title       string `json:"title" map:"title"`
	Description string `json:"description" map:"description"`
	Price       *Price `json:"price" map:"price"`
	Photo       string `json:"thumb_photo" map:"thumb_photo"`
	price       int
}

//func (p *Market) GetPrice() (price int) {
//	if p.price != 0 {
//		return p.price
//	}
//	p.price = pyraconv.ToFloat64(p.Price.Amount)
//	price = p.price
//	return
//}
