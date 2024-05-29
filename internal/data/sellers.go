package data

import (
	"context"
	"time"
)

const dbTimeout = time.Second * 3

type Seller struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (s Seller) GetAll() ([]Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, phone from sellers`

	//db := InitDB()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	// TODO catch err from rows.Close()
	defer rows.Close()

	var sellers []Seller

	for rows.Next() {
		var seller Seller
		if err := rows.Scan(
			&seller.ID,
			&seller.Name,
			&seller.Phone,
		); err != nil {
			return nil, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}
