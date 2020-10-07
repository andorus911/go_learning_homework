package interfaces

import (
	"context"
	"go_learning_homework/go-calendar-ms/internal/domain/models"
	"time"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, event *models.Event) error
	GetEventById(ctx context.Context, is string) (*models.Event, error)
	GetEventByOwnerStartTime(ctx context.Context, owner string, startTime time.Time) ([]*models.Event, error)
}
