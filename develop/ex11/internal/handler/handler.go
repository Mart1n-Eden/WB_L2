package handler

import (
	"encoding/json"
	"ex11/internal/handler/tools"
	"ex11/internal/model"
	"net/http"
	"time"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tools.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var ev model.Request

	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		tools.SendError(w, http.StatusBadRequest, "bad request")
		return
	}

	event, err := model.CastToEvent(ev)

	if err != nil {
		tools.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: error handling
	if err := h.service.Create(ev.UserId, event); err != nil {
		tools.SendError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	tools.SendSucsess(w, http.StatusOK, "successful creation")
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tools.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var ev model.Request

	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		tools.SendError(w, http.StatusBadRequest, "bad request")
		return
	}

	event, err := model.CastToEvent(ev)

	if err != nil {
		tools.SendError(w, http.StatusBadRequest, "invalid body")
	}

	if err := h.service.Update(ev.UserId, event); err != nil {
		switch err.Error() {
		case "past date":
			tools.SendError(w, http.StatusServiceUnavailable, err.Error())
		default:
			tools.SendError(w, http.StatusNotFound, "no id")
		}
		return
	}

	tools.SendSucsess(w, http.StatusOK, "successful update")
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tools.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := struct {
		ID int `json:"user_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		tools.SendError(w, http.StatusBadRequest, "bad request")
		return
	}

	if err := h.service.Delete(id.ID); err != nil {
		tools.SendError(w, http.StatusNotFound, "no id")
		return
	}

	tools.SendSucsess(w, http.StatusOK, "successful deletion")
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		tools.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// TODO:

	params := r.URL.Query()

	d, ok := params["date"]
	if !ok {
		tools.SendError(w, http.StatusBadRequest, "missing date")
		return
	}

	date, err := time.Parse("2006-01-02 15:04:05", d[0])

	if err != nil {
		tools.SendError(w, http.StatusBadRequest, "invalid date")
		return
	}

	var res []int

	if res, err = h.service.EventsForDay(date); err != nil {
		// TODO:
		tools.SendError(w, http.StatusBadRequest, "invalid date")
		return
	}

	tools.SendSucsess(w, http.StatusOK, res)
}

func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		tools.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// TODO:

	params := r.URL.Query()

	d, ok := params["date"]
	if !ok {
		tools.SendError(w, http.StatusBadRequest, "missing date")
		return
	}

	date, err := time.Parse("2006-01-02 15:04:05", d[0])

	if err != nil {
		tools.SendError(w, http.StatusBadRequest, "invalid date")
		return
	}

	var res []int

	if res, err = h.service.EventsForWeek(date); err != nil {
		// TODO:
		tools.SendError(w, http.StatusBadRequest, "invalid date")
		return
	}

	tools.SendSucsess(w, http.StatusOK, res)
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		tools.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// TODO:

	params := r.URL.Query()

	d, ok := params["date"]
	if !ok {
		tools.SendError(w, http.StatusBadRequest, "missing date")
		return
	}

	date, err := time.Parse("2006-01-02 15:04:05", d[0])

	if err != nil {
		tools.SendError(w, http.StatusBadRequest, "invalid date")
		return
	}

	var res []int

	if res, err = h.service.EventsForMonth(date); err != nil {
		// TODO:
		tools.SendError(w, http.StatusBadRequest, "invalid date")
		return
	}

	tools.SendSucsess(w, http.StatusOK, res)
}
