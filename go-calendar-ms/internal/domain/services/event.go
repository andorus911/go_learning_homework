package services

import (
	"context"
	"go_learning_homework/go-calendar-ms/internal/domain/models"
	"go_learning_homework/go-calendar-ms/internal/interfaces"
	"time"
)

type EventService struct {
	EventStorage interfaces.EventStorage
}

func (es *EventService) CreateEvent(ctx context.Context, title, description, owner string, startTime, endTime time.Time) (*models.Event, error) {
	event := &models.Event{
		Id:          "0", // todo uuid gen
		Title:       title,
		Description: description,
		Owner:       owner,
		StartTime:   startTime,
		EndTime:     endTime,
	}
	err := es.EventStorage.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (es *EventService) RemoveEvent(ctx context.Context, event *models.Event) error {
	err := es.EventStorage.DeleteEvent(ctx, event)
	return err
}

func (es *EventService) RemoveEventById(ctx context.Context, id string) error {
	event, err := es.EventStorage.GetEventById(ctx, id)
	if err != nil {
		return err
	}
	err = es.EventStorage.DeleteEvent(ctx, event)
	return err
}

func (es *EventService) EditEvent(ctx context.Context, id string, newEvent models.EditEvent) (*models.Event, error) {
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

	err = es.EventStorage.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (es *EventService) ShowOwnersEvents(ctx context.Context, owner string) ([]*models.Event, error) {
	events, err := es.EventStorage.GetEventByOwnerStartTime(ctx, owner, time.Now())
	if err != nil {
		return nil, err
	}
	return events, nil
}
