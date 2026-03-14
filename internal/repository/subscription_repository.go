package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"subscriptions/internal/model"
)

type SubscriptionRepository struct {
	db *sqlx.DB
}

func NewSubscriptionRepository(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(ctx context.Context, s *model.Subscription) error {
	query := `
	INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
	VALUES ($1,$2,$3,$4,$5) RETURNING id`
	return r.db.QueryRowContext(ctx, query, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate).Scan(&s.ID)
}

func (r *SubscriptionRepository) GetAll(ctx context.Context) ([]model.Subscription, error) {
	var subs []model.Subscription
	err := r.db.SelectContext(ctx, &subs, "SELECT * FROM subscriptions")
	return subs, err
}

func (r *SubscriptionRepository) GetByID(ctx context.Context, id int) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.GetContext(ctx, &sub, "SELECT * FROM subscriptions WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, s *model.Subscription) error {
	query := `
	UPDATE subscriptions
	SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5
	WHERE id=$6`
	_, err := r.db.ExecContext(ctx, query, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, s.ID)
	return err
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM subscriptions WHERE id=$1", id)
	return err
}

func (r *SubscriptionRepository) GetTotal(ctx context.Context, userID, serviceName, startDate, endDate string) (int, error) {
	query := "SELECT SUM(price) FROM subscriptions WHERE 1=1"
	args := []interface{}{}
	argID := 1

	if userID != "" {
		query += fmt.Sprintf(" AND user_id=$%d", argID)
		args = append(args, userID)
		argID++
	}
	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name=$%d", argID)
		args = append(args, serviceName)
		argID++
	}
	if startDate != "" {
		query += fmt.Sprintf(" AND start_date >= $%d", argID)
		args = append(args, startDate)
		argID++
	}
	if endDate != "" {
		query += fmt.Sprintf(" AND start_date <= $%d", argID)
		args = append(args, endDate)
		argID++
	}

	var total int
	err := r.db.GetContext(ctx, &total, query, args...)
	return total, err
}
