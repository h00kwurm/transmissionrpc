package transmissionrpc

type Tracker struct {
	Announce string `json:"announce, omitempty"`
	Id       string `json:"id"`
	Scrape   string `json:"scrape"`
	Tier     int    `json:"tier"`
}
