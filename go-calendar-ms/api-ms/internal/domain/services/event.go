package services

import (
	"context"
	"go_learning_homework/go-calendar-ms/api-ms/internal/domain/models"
	"go_learning_homework/go-calendar-ms/api-ms/internal/interfaces"
	"time"
)

type EventService struct {
	EventStorage interfaces.EventStorage
}

func (es *EventService) CreateEvent(ctx context.Context, title, description string, owner int64, startTime, endTime time.Time) (*models.Event, error) {
	event := &models.Event{
		Id:          0, // todo uuid gen
		Title:       title,
		Description: description,
		Owner:       owner,
		StartTime:   startTime,
		EndTime:     endTime,
	}
	id, err := es.EventStorage.SaveEvent(ctx, *event)
	if err != nil {
		return nil, err
	}
	event.Id = id
	return event, nil
}

func (es *EventService) RemoveEvent(ctx context.Context, event *models.Event) error {
	err := es.EventStorage.DeleteEventById(ctx, event.Id)
	return err
}

func (es *EventService) RemoveEventById(ctx context.Context, id int64) error {
	event, err := es.EventStorage.GetEventById(ctx, id)
	if err != nil {
		return err
	}
	err = es.EventStorage.DeleteEventById(ctx, event.Id)
	return err
}

func (es *EventService) EditEvent(ctx context.Context, id int64, newEvent models.EditEvent) (*models.Event, error) {
	event, err := es.EventStorage.GetEventById(ctx, id)
	if err != nil {
		return nil, err
	}

	if newEvent.Description != nil {
		event.Description = *newEvent.Description
	}
	if newEvent.Title != nil {
		event.Title = *newEvent.Title
	}
	if newEvent.StartTime != nil {
		event.StartTime = *newEvent.StartTime
	}
	if newEvent.EndTime != nil {
		event.EndTime = *newEvent.EndTime
	}

	id, err = es.EventStorage.SaveEvent(ctx, *event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (es *EventService) ShowOwnersEvents(ctx context.Context, owner int64) ([]models.Event, error) {
	events, err := es.EventStorage.GetEventsByOwnerStartTime(ctx, owner, time.Now())
	if err != nil {
		return nil, err
	}
	return events, nil
}
