package api

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"go_learning_homework/go-calendar-ms/internal/domain/models"
	"go_learning_homework/go-calendar-ms/internal/domain/services"
	"net/http"
	"strconv"
	"time"
)

type respEvent struct {
	Data models.Event `json:"data"`
}
type respEventList struct {
	Items []models.Event `json:"items"`
	Count int `json:"count"`
}
type respInfo struct {
	Info string `json:"info"`
}

const LAYOUT = "2006.01.02 15:04:05"

var lg zap.Logger
var evService services.EventService

func StartServer(httpListen string, logger zap.Logger, eventService *services.EventService) {
	lg = logger
	evService = *eventService

	evh := &EventHandler{}
	mux := http.NewServeMux()
	mux.HandleFunc("/event/create", evh.create) // timeout?
	mux.HandleFunc("/event/update", evh.update)
	mux.HandleFunc("/event/delete", evh.delete)
	mux.HandleFunc("/event/myList", evh.getList)

	server := http.Server{
		Addr: httpListen,
		Handler: mux,
	}

	//go func() {
		err := server.ListenAndServe()
		if err != nil {
			lg.Fatal(err.Error())
		}
	//}()
}

type EventHandler struct{}

func (h *EventHandler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		lg.Warn("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Allow: POST"))
		if err != nil {
			lg.Error("response writer error: " + err.Error())
		}
		return
	}

	title := r.PostFormValue("title")
	if title == "" {
		lg.Warn("incorrect title")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	description := r.PostFormValue("description")
	owner, err := strconv.ParseInt(r.PostFormValue("owner"), 10, 64)
	if err != nil {
		lg.Warn("incorrect owner: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	startTime, err := time.Parse(LAYOUT, r.PostFormValue("startTime"))
	if err != nil {
		lg.Warn("incorrect startTime: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse(LAYOUT, r.PostFormValue("endTime"))
	if err != nil {
		lg.Warn("incorrect endTime: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	event, err := evService.CreateEvent(ctx, title, description, owner, startTime, endTime)
	if err != nil {
		lg.Error("event creating error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(respEvent{Data: *event})
	_, err = w.Write(out)
	if err != nil {
		lg.Error("response writer error: " + err.Error())
	}
	return
}

func (h *EventHandler) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" && r.Method != "PUT" {
		lg.Warn("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Allow: POST"))
		if err != nil {
			lg.Error("response writer error: " + err.Error())
		}
		return
	}

	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		lg.Warn("incorrect owner: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	editEvent := models.EditEvent{}
	if title := r.PostFormValue("title"); title != "" {
		editEvent.Title = &title
	}
	if description := r.PostFormValue("description"); description != "" {
		editEvent.Title = &description
	}
	if r.PostFormValue("startTime") != "" {
		startTime, err := time.Parse(LAYOUT, r.PostFormValue("endTime"))
		if err != nil {
			lg.Warn("incorrect startTime: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		editEvent.StartTime = &startTime
	}

	if r.PostFormValue("endTime") != "" {
		endTime, err := time.Parse(LAYOUT, r.PostFormValue("endTime"))
		if err != nil {
			lg.Warn("incorrect endTime: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		editEvent.EndTime = &endTime
	}

	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	event, err := evService.EditEvent(ctx, id, editEvent)
	if err != nil {
		lg.Error("event editing error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	out, err := json.Marshal(respEvent{Data: *event})
	_, err = w.Write(out)
	if err != nil {
		lg.Error("response writer error: " + err.Error())
	}
	return
}

func (h *EventHandler) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" && r.Method != "DELETE" {
		lg.Warn("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Allow: POST, DELETE"))
		if err != nil {
			lg.Error("response writer error: " + err.Error())
		}
		return
	}

	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		lg.Warn("incorrect owner: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	err = evService.RemoveEventById(ctx, id)
	if err != nil {
		lg.Error("event editing error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	out, err := json.Marshal(respInfo{Info: "success"})
	_, err = w.Write(out)
	if err != nil {
		lg.Error("response writer error: " + err.Error())
	}
	return
}

func (h *EventHandler) getList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		lg.Warn("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Allow: GET"))
		if err != nil {
			lg.Error("response writer error: " + err.Error())
		}
		return
	}

	owner, err := strconv.ParseInt(r.PostFormValue("owner"), 10, 64)
	if err != nil {
		lg.Warn("incorrect owner: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx, _ := context.WithTimeout(context.Background(), 555 * time.Second)
	events, err := evService.ShowOwnersEvents(ctx, owner)
	if err != nil {
		lg.Error("event editing error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	out, err := json.Marshal(respEventList{
		Items: events,
		Count: len(events),
	})
	_, err = w.Write(out)
	if err != nil {
		lg.Error("response writer error: " + err.Error())
	}
	return
}