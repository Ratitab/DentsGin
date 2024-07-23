package Dents

type Dent struct {
	ID    string `json:"_id"  bson:"_id"`
	Name  string `json:"username" bson:"username"`
	Email string `json:"password" bson:"password"`
}

type Treatment struct {
	Disease  string  `json:"disease"`
	Text     string  `json:"text"`
	Quantity int     `json:"quantity"`
	OnePrice float64 `json:"onePrice"`
	Total    float64 `json:"total"`
}

type Phase struct {
	ID           int64       `json:"id"`
	ClickedTeeth []int       `json:"clickedTeeth"`
	Days         string      `json:"days"`
	Treatments   []Treatment `json:"treatments"`
}

type PacientData struct {
	Email  string  `json:"email"`
	Name   string  `json:"name"`
	Phases []Phase `json:"phases"`
}

type SearchItem struct {
	Name string `json:"name" bson:"name"`
}

type PaymentStatus struct {
	IsPaid bool `json:isPaid bson:"isPaid"`
}

type VersionResponse struct {
	Version     string `json:"version" bson:"version"`
	DownloadURL string `json:"downloadURL" bson:"downloadURL"`
}
