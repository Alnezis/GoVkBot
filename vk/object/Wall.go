package object

type Wall struct {
	Id      float64 `json:"id" map:"id"`
	FromId  float64 `json:"from_id" map:"from_id"`
	OwnerId float64 `json:"owner_id" map:"owner_id"`
}
