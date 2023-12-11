package Manufacturers

import "time"

type Manufacturer struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Logo      *string    `json:"logo"`
	TypeID    *int       `json:"type_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
