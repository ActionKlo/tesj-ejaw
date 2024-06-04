package repository

import (
	"context"
	"database/sql"
	"github.com/ActionKlo/test-ejaw/internal/models"
	"go.uber.org/zap"
	"time"
)

const dbTimeout = time.Second * 3

func (s *Service) GetAll() ([]models.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, phone from sellers`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			s.logger.Error("failed to close sql.Rows:", zap.Error(err))
		}
	}(rows)

	var sellers []models.Seller

	for rows.Next() {
		var seller models.Seller
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

func (s *Service) Insert(name, phone string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `insert into sellers (name, phone) values ($1, $2) returning id`

	var newSellerID int
	if err := s.db.QueryRowContext(ctx, query,
		name,
		phone,
	).Scan(&newSellerID); err != nil {
		return 0, err
	}

	return newSellerID, nil
}

func (s *Service) Delete(id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `delete from sellers where id = $1`

	res, err := s.db.ExecContext(ctx, query, id)
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

func (s *Service) Update(id int, name, phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `update sellers set
				name = $2,
				phone = $3
				where id = $1
	`

	res, err := s.db.ExecContext(ctx, query,
		id,
		name,
		phone,
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
