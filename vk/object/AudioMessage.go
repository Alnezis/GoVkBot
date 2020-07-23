package object

type AudioMessage struct {
	Duration int    `json:"duration" map:"duration"`
	Waveform []int  `json:"waveform" map:"waveform"`
	LinkOGG  string `json:"link_ogg" map:"link_ogg"`
	LinkMP3  string `json:"link_mp3" map:"link_mp3"`
}
