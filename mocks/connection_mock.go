package mocks

import (
	"context"

	"github.com/redundant4u/DeoDeokGo/internal/models"
	"github.com/redundant4u/DeoDeokGo/internal/queue"
)

type MockEventEmitter struct{}

type MockEventsRepository struct{}

func (r *MockEventsRepository) Add(ctx context.Context, e models.Event) ([]byte, error) {
	return nil, nil
}

func (r *MockEventsRepository) Find(ctx context.Context, id string) (models.Event, error) {
	return models.Event{}, nil
}

func (r *MockEventsRepository) FindByName(ctx context.Context, name string) (models.Event, error) {
	return models.Event{}, nil
}

func (r *MockEventsRepository) FindAll(ctx context.Context) ([]models.Event, error) {
	return []models.Event{}, nil
}

func (e *MockEventEmitter) Emit(event queue.Event) error {
	return nil
}
