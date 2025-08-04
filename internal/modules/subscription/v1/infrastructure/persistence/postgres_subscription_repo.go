package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/modules/subscription/v1/domain/repo"
	"subscriptions/internal/shared/persistence"
	"time"
)

type repository struct {
	DB *persistence.Postgres
}

func NewRepository(db *persistence.Postgres) repo.Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(ctx context.Context, model *model.Subscription) (*model.Subscription, error) {
	query := `
		INSERT INTO subscriptions (name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.DB.QueryRowContext(ctx, query,
		model.Name,
		model.Price,
		model.UserId,
		model.StartDate,
		model.EndDate,
	).Scan(&model.ID)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*model.Subscription, error) {
	query := `
		SELECT id, name, price, user_id, start_date, end_date
		FROM subscriptions
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}(rows)

	var subscriptions []*model.Subscription
	for rows.Next() {
		var sub model.Subscription
		if err := rows.Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserId, &sub.StartDate, &sub.EndDate); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, &sub)
	}

	return subscriptions, nil
}

func (r *repository) GetById(ctx context.Context, id uint) (*model.Subscription, error) {
	query := `
		SELECT id, name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE id = $1
	`

	var sub model.Subscription
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserId, &sub.StartDate, &sub.EndDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &sub, nil
}

func (r *repository) Update(ctx context.Context, model *model.Subscription) (*model.Subscription, error) {
	query := `
		UPDATE subscriptions
		SET name = $1, price = $2, user_id = $3, start_date = $4, end_date = $5
		WHERE id = $6
	`

	_, err := r.DB.ExecContext(ctx, query,
		model.Name,
		model.Price,
		model.UserId,
		model.StartDate,
		model.EndDate,
		model.ID,
	)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	query := `
		DELETE FROM subscriptions
		WHERE id = $1
	`

	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *repository) CalculateTotalCost(ctx context.Context, startDate, endDate time.Time, userId *uuid.UUID, subscriptionName *string) (float64, error) {
	query := `
		SELECT SUM(price)
		FROM subscriptions
		WHERE start_date >= $1 AND (end_date <= $2 OR end_date IS NULL)
	`

	args := []interface{}{startDate, endDate}
	argIndex := 3

	if userId != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argIndex)
		args = append(args, *userId)
		argIndex++
	}

	if subscriptionName != nil {
		query += fmt.Sprintf(" AND name = $%d", argIndex)
		args = append(args, *subscriptionName)
	}

	var totalCost sql.NullFloat64
	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&totalCost)
	if err != nil {
		return 0, err
	}

	if !totalCost.Valid {
		return 0, nil
	}

	return totalCost.Float64, nil
}
