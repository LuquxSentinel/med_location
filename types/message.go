package types

type Message struct {
	Channel  string   `json:"channel"`
	Location Location `json:"location"`
}
