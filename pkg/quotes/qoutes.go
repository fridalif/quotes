package quotes

type Quote struct {
	Id     uint   `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}
