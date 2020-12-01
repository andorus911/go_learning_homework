package dbevent

import (
	"context"
	"errors"
	"go_learning_homework/go-calendar-ms/api-ms/internal/domain/models"
	"time"
)

type DBtype map[string]models.Event

func InitDB() {
	//DB := make(map[string]models.Event)
}

func (DB DBtype) SaveEvent(ctx context.Context, event *models.Event) error {
	DB[event.Id] = *event
	if DB[event.Id] != *event {
		return errors.New("db error")
	}
	return nil
}

func (DB DBtype) DeleteEvent(ctx context.Context, event *models.Event) error {
	delete(DB, event.Id)
	return nil
}

func (DB DBtype) GetEventById(ctx context.Context, id string) (models.Event, error) {
	return DB[id], nil
}

func (DB DBtype) GetEventByOwnerStartTime(ctx context.Context, owner int, startTime time.Time) ([]models.Event, error) {
	events := make([]models.Event, 1)
	for _, event := range DB {
		if event.Owner == owner &&
			(event.StartTime.After(startTime) || event.StartTime.Equal(startTime)) {
			events = append(events, event)
		}
	}
	return events, nil
}