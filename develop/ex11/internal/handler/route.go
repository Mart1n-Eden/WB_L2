package handler

import (
	"ex11/internal/handler/middleware"
	"ex11/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/create_event", middleware.Logging(h.Create))
	r.HandleFunc("/update_event", middleware.Logging(h.Update))
	r.HandleFunc("/delete_event", middleware.Logging(h.Delete))
	r.HandleFunc("/events_for_day", middleware.Logging(h.EventsForDay))
	r.HandleFunc("/events_for_week", middleware.Logging(h.EventsForWeek))
	r.HandleFunc("/events_for_month", middleware.Logging(h.EventsForMonth))

	return r
}

//POST /create_event
//POST /update_event
//POST /delete_event
//GET /events_for_day
//GET /events_for_week
//GET /events_for_month
