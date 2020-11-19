package http2server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	gc "go_learning_homework/go-calendar-ms/api/grpc"
	"go_learning_homework/go-calendar-ms/internal/domain/models"
	"go_learning_homework/go-calendar-ms/internal/domain/services"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
)

var lg zap.Logger
var evService services.EventService

func RunGrpcServer(zlg zap.Logger, address, port string, eventService services.EventService) {
	evService = eventService
	lg = zlg
	lis, err := net.Listen("tcp", address + ":" + port)
	if err != nil {
		lg.Fatal("incorrect server start : " + err.Error())
	}

	grpcServer := grpc.NewServer()
	gc.RegisterGoCalendarCRUDServer(grpcServer, ServerDull{})
	err = grpcServer.Serve(lis)
	if err != nil {
		lg.Error("server was stopped : " + err.Error())
	}
}

type ServerDull struct {
	gc.UnimplementedGoCalendarCRUDServer
}

func (sd ServerDull) CreateEvent(ctx context.Context, message *gc.CreateEventMessage) (*gc.EventResponseMessage, error) {
	event, err := evService.CreateEvent(ctx, message.Title, message.Description, message.Owner, message.StartTime.AsTime(), message.EndTime.AsTime())
	if err != nil {
		lg.Error("event creating error: " + err.Error())
		return nil, err
	}

	return &gc.EventResponseMessage{
		Id:          event.Id,
		Title:       event.Title,
		Description: event.Description,
		Owner:       event.Owner,
		StartTime:   timestamppb.New(event.StartTime),
		EndTime:     timestamppb.New(event.EndTime),
	}, nil
}

func (sd ServerDull) ReadEvents(ctx context.Context, message *gc.ReadEventsMessage) (*gc.EventsResponseMessage, error) {
	events, err := evService.ShowOwnersEvents(ctx, message.Owner)
	if err != nil {
		lg.Error("event editing error: " + err.Error())
		return nil, err
	}

	var eventsRes gc.EventsResponseMessage
	for _, e := range events {
		er := gc.EventResponseMessage{
			Id:          e.Id,
			Title:       e.Title,
			Description: e.Description,
			Owner:       e.Owner,
			StartTime:   timestamppb.New(e.StartTime),
			EndTime:     timestamppb.New(e.EndTime),
		}
		eventsRes.Events = append(eventsRes.Events, &er)
	}
	return &eventsRes, nil
}

func (sd ServerDull) UpdateEvent(ctx context.Context, message *gc.UpdateEventMessage) (*gc.EventResponseMessage, error) {
	ee := models.EditEvent{
		Title:       message.Title,
		Description: message.Description,
	}

	if message.StartTime != nil {
		t := message.StartTime.AsTime()
		ee.StartTime = &t
	} else {
		ee.StartTime = nil
	}

	if message.EndTime != nil {
		t := message.EndTime.AsTime()
		ee.EndTime = &t
	} else {
		ee.EndTime = nil
	}

	event, err := evService.EditEvent(ctx, message.Id, ee)
	if err != nil {
		lg.Error("event editing error: " + err.Error())
		return nil, err
	}

	return &gc.EventResponseMessage{
		Id:          event.Id,
		Title:       event.Title,
		Description: event.Description,
		Owner:       event.Owner,
		StartTime:   timestamppb.New(event.StartTime),
		EndTime:     timestamppb.New(event.EndTime),
	}, nil
}

func (sd ServerDull) DeleteEvent(ctx context.Context, message *gc.DeleteEventMessage) (*empty.Empty, error) {
	err := evService.RemoveEventById(ctx, message.Id)
	if err != nil {
		lg.Error("event editing error: " + err.Error())
		return nil, err
	}

	return nil, nil
}
