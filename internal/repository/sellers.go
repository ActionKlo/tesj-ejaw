package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

type Seller struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func (s Seller) GetAll() ([]Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, phone from sellers`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Println("failed to close sql.Rows:", err.Error())
		}
	}(rows)

	var sellers []Seller

	for rows.Next() {
		var seller Seller
		if err = rows.Scan(
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

func (s Seller) Insert() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `insert into sellers (name, phone) values ($1, $2) returning id`

	var newSellerID int
	if err := db.QueryRowContext(ctx, query,
		s.Name,
		s.Phone,
	).Scan(&newSellerID); err != nil {
		return 0, err
	}

	return newSellerID, nil
}

func (s Seller) Delete() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `delete from sellers where id = $1`

	res, err := db.ExecContext(ctx, query, s.ID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (s Seller) Update() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `update sellers set
				name = $2,
				phone = $3
				where id = $1
	`

	res, err := db.ExecContext(ctx, query,
		s.ID,
		s.Name,
		s.Phone,
	)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}