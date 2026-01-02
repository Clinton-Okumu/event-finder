package api

import (
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type TicketsHandler struct {
	bookingStore     store.BookingStore
	bookingItemStore store.BookingItemStore
	eventTicketStore store.EventTicketStore
	eventStore       store.EventStore
	logger           *log.Logger
}

func NewTicketsHandler(bookingStore store.BookingStore, bookingItemStore store.BookingItemStore, eventTicketStore store.EventTicketStore, eventStore store.EventStore, logger *log.Logger) *TicketsHandler {
	return &TicketsHandler{
		bookingStore:     bookingStore,
		bookingItemStore: bookingItemStore,
		eventTicketStore: eventTicketStore,
		eventStore:       eventStore,
		logger:           logger,
	}
}

type TicketResponse struct {
	ID            uint    `json:"id"`
	UserID        uint    `json:"user_id"`
	EventID       uint    `json:"event_id"`
	EventTitle    string  `json:"event_title"`
	EventDate     string  `json:"event_date"`
	EventLocation string  `json:"event_location"`
	EventImageURL string  `json:"event_image_url"`
	PricePaid     float64 `json:"price_paid"`
	BookingDate   string  `json:"booking_date"`
	QRCode        string  `json:"qr_code,omitempty"`
}

type BookTicketRequest struct {
	EventID      uint `json:"event_id"`
	TicketTypeID uint `json:"ticket_type_id"`
	Quantity     int  `json:"quantity"`
}

func (h *TicketsHandler) GetUserTickets(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	if user.IsAnonymous() {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "You must be logged in"})
		return
	}

	bookings, err := h.bookingStore.GetBookingsByUserID(r.Context(), user.ID)
	if err != nil {
		h.logger.Printf("ERROR: fetching bookings: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve tickets"})
		return
	}

	tickets := make([]TicketResponse, 0, len(bookings))
	for _, booking := range bookings {
		if booking.Status == models.BookingCanceled {
			continue
		}

		totalPrice := 0.0
		if len(booking.BookingItems) > 0 {
			for _, item := range booking.BookingItems {
				totalPrice += float64(item.Quantity) * item.PriceAtPurchase
			}
		} else if booking.Event != nil {
			totalPrice = booking.Event.Price
		}

		eventDate := ""
		if booking.Event != nil {
			eventDate = booking.Event.StartTime.Format("2006-01-02 15:04:05")
		}

		ticket := TicketResponse{
			ID:            booking.ID,
			UserID:        booking.UserID,
			EventID:       booking.EventID,
			EventTitle:    booking.Event.Title,
			EventDate:     eventDate,
			EventLocation: booking.Event.Location,
			EventImageURL: booking.Event.ImageURL,
			PricePaid:     totalPrice,
			BookingDate:   booking.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		tickets = append(tickets, ticket)
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"tickets": tickets})
}

func (h *TicketsHandler) BookTicket(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	if user.IsAnonymous() {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "You must be logged in"})
		return
	}

	var req BookTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Printf("ERROR: decoding book ticket request: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	if req.EventID == 0 {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "event_id is required"})
		return
	}

	event, err := h.eventStore.GetEvent(r.Context(), req.EventID)
	if err != nil {
		h.logger.Printf("ERROR: fetching event: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "event not found"})
		return
	}

	if event.TicketsRemaining <= 0 {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "no tickets available for this event"})
		return
	}

	ticketType := &models.EventTicket{}
	if req.TicketTypeID != 0 {
		ticketType, err = h.eventTicketStore.GetEventTicket(r.Context(), req.TicketTypeID)
		if err != nil {
			h.logger.Printf("ERROR: fetching ticket type: %v", err)
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "ticket type not found"})
			return
		}
	}

	quantity := req.Quantity
	if quantity <= 0 {
		quantity = 1
	}

	if ticketType.ID != 0 && ticketType.Quantity < quantity {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "not enough tickets available"})
		return
	}

	booking := &models.Booking{
		UserID:  user.ID,
		EventID: req.EventID,
		Status:  models.BookingConfirmed,
	}

	if err := h.bookingStore.CreateBooking(r.Context(), booking); err != nil {
		h.logger.Printf("ERROR: creating booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create booking"})
		return
	}

	totalPrice := event.Price

	if ticketType.ID != 0 {
		bookingItem := &models.BookingItem{
			BookingID:       booking.ID,
			TicketTypeID:    ticketType.ID,
			Quantity:        quantity,
			PriceAtPurchase: ticketType.Price,
		}

		if err := h.bookingItemStore.CreateBookingItem(r.Context(), bookingItem); err != nil {
			h.logger.Printf("ERROR: creating booking item: %v", err)
		}
		totalPrice = ticketType.Price * float64(quantity)
	}

	response := TicketResponse{
		ID:            booking.ID,
		UserID:        booking.UserID,
		EventID:       booking.EventID,
		EventTitle:    event.Title,
		EventDate:     event.StartTime.Format("2006-01-02 15:04:05"),
		EventLocation: event.Location,
		EventImageURL: event.ImageURL,
		PricePaid:     totalPrice,
		BookingDate:   booking.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"ticket": response})
}

func (h *TicketsHandler) CancelTicket(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	if user.IsAnonymous() {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "You must be logged in"})
		return
	}

	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid ticket id"})
		return
	}

	booking, err := h.bookingStore.GetBooking(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrBookingNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "ticket not found"})
			return
		}
		h.logger.Printf("ERROR: fetching booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve ticket"})
		return
	}

	if booking.UserID != user.ID {
		utils.WriteJSON(w, http.StatusForbidden, utils.Envelope{"error": "you can only cancel your own tickets"})
		return
	}

	if err := h.bookingStore.DeleteBooking(r.Context(), uint(id)); err != nil {
		h.logger.Printf("ERROR: deleting booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not cancel ticket"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "ticket cancelled successfully"})
}
