package Pacients

type Pacient struct {
	Email  string  `json:"email"`
	Name   string  `json:"name"`
	Phases []Phase `json:"phases"`
}
