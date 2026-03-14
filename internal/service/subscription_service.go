package service

import (
	"context"
	"subscriptions/internal/model"
	"subscriptions/internal/repository"
)

type SubscriptionService struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionService(repo *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(ctx context.Context, sub *model.Subscription) error {
	return s.repo.Create(ctx, sub)
}

func (s *SubscriptionService) GetAll(ctx context.Context) ([]model.Subscription, error) {
	return s.repo.GetAll(ctx)
}

func (s *SubscriptionService) GetByID(ctx context.Context, id int) (*model.Subscription, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SubscriptionService) Update(ctx context.Context, sub *model.Subscription) error {
	return s.repo.Update(ctx, sub)
}

func (s *SubscriptionService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *SubscriptionService) GetTotal(ctx context.Context, userID, serviceName, startDate, endDate string) (int, error) {
	return s.repo.GetTotal(ctx, userID, serviceName, startDate, endDate)
}
