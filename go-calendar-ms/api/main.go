package api

import (
	"go.uber.org/zap"
	"log"
	"net/http"
)

var lg zap.Logger

func StartServer(httpListen string, logger zap.Logger) {
	lg = logger

	server := http.Server{
		Addr: httpListen,
		Handler: nil, // TODO
	}

	//go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	//}()
}

type EventHandler struct{}

func (h *EventHandler) create(w http.ResponseWriter, r *http.Request) {
	//r.PostFormValue()
	//services.EventService{}
}

func (h *EventHandler) update(w http.ResponseWriter, r http.Request) {

}

func (h *EventHandler) delete(w http.ResponseWriter, r http.Request) {

}

func (h *EventHandler) getList(w http.ResponseWriter, r http.Request) {

}