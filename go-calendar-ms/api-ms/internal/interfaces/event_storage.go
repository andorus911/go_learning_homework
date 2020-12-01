package interfaces

import (
	"context"
	"go_learning_homework/go-calendar-ms/api-ms/internal/domain/models"
	"time"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event models.Event) (int64, error)
	DeleteEventById(ctx context.Context, id int64) error
	GetEventById(ctx context.Context, id int64) (*models.Event, error)
	GetEventsByOwnerStartTime(ctx context.Context, owner int64, startTime time.Time) ([]models.Event, error)
}
