package Pacients

type Treatment struct {
	Disease  string  `json:"disease"`
	Text     string  `json:"text"`
	Quantity int     `json:"quantity"`
	OnePrice float64 `json:"onePrice"`
	Total    float64 `json:"total"`
}
