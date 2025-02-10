package handler

const JsonPath = "./json/"

type Item struct {
	Item                string `json:"item"`
	BPItem              string `json:"bpitem"`
	CoinsItem           string `json:"coinsItem"`
	CoinsItemBP         string `json:"coinsItemBP"`
	CoinsItemDesconto   string `json:"coinsItemDesconto"`
	CoinsItemBPDesconto string `json:"coinsItemBPDesconto"`
}

type Dino struct {
	Dino     string `json:"dino"`
	BP       string `json:"bp"`
	Coins    string `json:"coins"`
	Desconto string `json:"desconto"`
	Tokens   string `json:"tokens"`
}

type Chibi struct {
	Chibi   string `json:"chibi"`
	BPChibi string `json:"bpchibi"`
}
