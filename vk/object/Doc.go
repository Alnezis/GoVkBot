package object

import "fmt"

type Preview struct {
	Photo        *Photo        `json:"photo" map:"photo"`
	Graffiti     *Graffiti     `json:"graffiti" map:"graffiti"`
	AudioMessage *AudioMessage `json:"audio_message" map:"audio_message"`
}

type Doc struct {
	ID      int      `json:"id" map:"id"`
	OwnerID int      `json:"owner_id" map:"owner_id"`
	Title   string   `json:"title" map:"title"`
	Size    int      `json:"size" map:"size"`
	Ext     string   `json:"ext" map:"ext"`
	Url     string   `json:"url" map:"url"`
	Date    int      `json:"date" map:"date"`
	Type    int      `json:"type" map:"type"`
	Preview *Preview `json:"preview" map:"preview"`
}

func (d *Doc) BuildAttachment() string {
	return fmt.Sprintf("doc%d_%d", int(d.OwnerID), int(d.ID))
}
