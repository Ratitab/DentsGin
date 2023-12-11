package services

import (
	"gitlab.com/golanggin/initial/shadow/internal/models/Manufacturers"
	"gitlab.com/golanggin/initial/shadow/pkg/database/db_drivers/mysql"
)

type ManufacturerService struct {
	MySQL *mysql.MySQL // MySQL connection instance
}

func NewManufacturerService(mysql *mysql.MySQL) *ManufacturerService {
	return &ManufacturerService{
		MySQL: mysql,
	}
}

func (s *ManufacturerService) GetManufacturers() ([]Manufacturers.Manufacturer, error) {
	var manufacturers []Manufacturers.Manufacturer

	query := "SELECT id, title, logo, type_id, created_at, updated_at, deleted_at FROM manufacturers"
	rows, err := s.MySQL.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var manufacturer Manufacturers.Manufacturer
		if err := rows.Scan(
			&manufacturer.ID,
			&manufacturer.Title,
			&manufacturer.Logo,
			&manufacturer.TypeID,
			&manufacturer.CreatedAt,
			&manufacturer.UpdatedAt,
			&manufacturer.DeletedAt,
		); err != nil {
			return nil, err
		}

		manufacturers = append(manufacturers, manufacturer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return manufacturers, nil
}
