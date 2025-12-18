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
)

type EventTicketsHandler struct {
	eventTicketStore store.EventTicketStore
	logger           *log.Logger
}

func NewEventsTicketsHandler(store store.EventTicketStore, logger *log.Logger) *EventTicketsHandler {
	return &EventTicketsHandler{
		eventTicketStore: store,
		logger:           logger,
	}
}

// CreateEventTicket creates a new event ticket
func (h *EventTicketsHandler) CreateEventTicket(w http.ResponseWriter, r *http.Request) {
	var eventTicket models.EventTicket
	if err := json.NewDecoder(r.Body).Decode(&eventTicket); err != nil {
		h.logger.Printf("ERROR: decoding event ticket: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	if err := h.eventTicketStore.CreateEventTicket(r.Context(), &eventTicket); err != nil {
		h.logger.Printf("ERROR: creating event ticket: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create event ticket"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"event ticket": eventTicket})
}

// GetEventTickets fetches all event tickets with optional sorting and pagination
func (h *EventTicketsHandler) GetEventTickets(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))
	orderBy := q.Get("order_by")
	if orderBy == "" {
		orderBy = "created_at asc"
	}

	eventTickets, err := h.eventTicketStore.GetEventTickets(r.Context(), limit, offset, orderBy)
	if err != nil {
		h.logger.Printf("ERROR: fetching event tickets: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve event tickets"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"event tickets": eventTickets})

}

// GetEventTicketByID fetches a single event ticket
func (h *EventTicketsHandler) GetEventTicketByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid event ticket id"})
		return
	}

	eventTicket, err := h.eventTicketStore.GetEventTicket(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrEventNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "event ticket not found"})
			return
		}
		h.logger.Printf("ERROR: fetching event ticket: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve event ticket"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"event ticket": eventTicket})
}

// UpdateEventTicket updates an existing event ticket
func (h *EventTicketsHandler) UpdateEventTicket(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid event ticket id"})
		return
	}

	var eventTicket models.EventTicket
	if err := json.NewDecoder(r.Body).Decode(&eventTicket); err != nil {
		h.logger.Printf("ERROR: decoding event ticket: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	eventTicket.ID = uint(id)
	if err := h.eventTicketStore.UpdateEventTicket(r.Context(), &eventTicket); err != nil {
		if errors.Is(err, store.ErrEventTicketNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "event ticket not found"})
			return
		}
		h.logger.Printf("ERROR: updating event ticket: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update event ticket"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"event ticket": eventTicket})

}

// DeleteEventTicket deletes a event ticket by ID
func (h *EventTicketsHandler) DeleteEventTicket(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid event ticket id"})
		return
	}
	if err := h.eventTicketStore.DeleteEventTicket(r.Context(), uint(id)); err != nil {
		if errors.Is(err, store.ErrEventTicketNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "event ticket not found"})
			return
		}
		h.logger.Printf("ERROR: deleting event ticket: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete event ticket"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "event ticket deleted"})

}
