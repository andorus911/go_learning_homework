package api

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"go_learning_homework/go-calendar-ms/internal/domain/services"
	"net/http"
	"strconv"
	"time"
)

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
	}
	description := r.PostFormValue("description")
	owner, err := strconv.ParseInt(r.PostFormValue("owner"), 10, 64)
	if err != nil {
		lg.Warn("incorrect owner: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	startTime, err := time.Parse(LAYOUT, r.PostFormValue("startTime"))
	if err != nil {
		lg.Warn("incorrect startTime: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	endTime, err := time.Parse(LAYOUT, r.PostFormValue("endTime"))
	if err != nil {
		lg.Warn("incorrect endTime: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	event, err := evService.CreateEvent(ctx, title, description, owner, startTime, endTime)
	if err != nil {
		lg.Error("event creating error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	out, err := json.Marshal(event)
	_, err = w.Write(out)
	if err != nil {
		lg.Error("response writer error: " + err.Error())
	}
	return
}

func (h *EventHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *EventHandler) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *EventHandler) getList(w http.ResponseWriter, r *http.Request) {

}