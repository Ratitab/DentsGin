package Pacients

type Phase struct {
	ID           int64       `json:"id"`
	ClickedTeeth []int       `json:"clickedTeeth"`
	Days         string      `json:"days"`
	Treatments   []Treatment `json:"treatments"`
}
