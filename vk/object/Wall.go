package object

type Wall struct {
	Id      int `json:"id" map:"id"`
	FromId  int `json:"from_id" map:"from_id"`
	OwnerId int `json:"owner_id" map:"owner_id"`
}
