package api

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type EventsHandler struct {
	eventStore store.EventStore
	logger     *log.Logger
}

func NewEventsHandler(store store.EventStore, logger *log.Logger) *EventsHandler {
	return &EventsHandler{
		eventStore: store,
		logger:     logger,
	}
}

// CreateEvent creates a new event
func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.logger.Printf("ERROR: decoding event: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	if err := h.eventStore.CreateEvent(r.Context(), &event); err != nil {
		h.logger.Printf("ERROR: creating event: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create event"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"event": event})
}

// GetEvents fetches all events with optional sorting and pagination
func (h *EventsHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))
	orderBy := q.Get("order_by")
	if orderBy == "" {
		orderBy = "created_at asc" // default sorting
	}

	events, err := h.eventStore.GetEvents(r.Context(), limit, offset, orderBy)
	if err != nil {
		h.logger.Printf("ERROR: fetching events: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve events"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"events": events})
}

// GetEventByID fetches a single event
func (h *EventsHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid event id"})
		return
	}

	event, err := h.eventStore.GetEvent(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrEventNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "event not found"})
			return
		}
		h.logger.Printf("ERROR: fetching event: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve event"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"event": event})
}

// UpdateEvent updates an existing event
func (h *EventsHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid event id"})
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.logger.Printf("ERROR: decoding event: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	// Basic validation
	event.Title = strings.TrimSpace(event.Title)
	if event.Title == "" {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "event title is required"})
		return
	}

	event.ID = uint(id)

	if err := h.eventStore.UpdateEvent(r.Context(), &event); err != nil {
		if errors.Is(err, store.ErrEventNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "event not found"})
			return
		}
		h.logger.Printf("ERROR: updating event: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update event"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"event": event})
}

// DeleteEvent deletes a event by ID
func (h *EventsHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid event id"})
		return
	}

	if err := h.eventStore.DeleteEvent(r.Context(), uint(id)); err != nil {
		if errors.Is(err, store.ErrEventNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "event not found"})
			return
		}
		h.logger.Printf("ERROR: deleting event: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete event"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "event deleted"})
}
