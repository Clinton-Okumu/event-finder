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

type BookingsHandler struct {
	bookingStore store.BookingStore
	logger       *log.Logger
}

func NewBookingsHandler(store store.BookingStore, logger *log.Logger) *BookingsHandler {
	return &BookingsHandler{
		bookingStore: store,
		logger:       logger,
	}
}

// CreateBooking creates a new booking
func (h *BookingsHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		h.logger.Printf("ERROR: decoding booking: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	if err := h.bookingStore.CreateBooking(r.Context(), &booking); err != nil {
		h.logger.Printf("ERROR: creating booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create booking"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"booking": booking})
}

// GetBookings fetches all bookings with optional sorting and pagination
func (h *BookingsHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))
	orderBy := q.Get("order_by")
	if orderBy == "" {
		orderBy = "created_at asc" // default sorting
	}

	bookings, err := h.bookingStore.GetBookings(r.Context(), limit, offset, orderBy)
	if err != nil {
		h.logger.Printf("ERROR: fetching bookings: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve bookings"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"bookings": bookings})

}

// GetBookingByID fetches a single booking
func (h *BookingsHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking id"})
		return
	}

	booking, err := h.bookingStore.GetBooking(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrBookingNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "booking not found"})
			return
		}
		h.logger.Printf("ERROR: fetching booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve booking"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"booking": booking})
}

// UpdateBooking updates an existing booking
func (h *BookingsHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking id"})
		return
	}

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		h.logger.Printf("ERROR: decoding booking: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	booking.ID = uint(id)
	if err := h.bookingStore.UpdateBooking(r.Context(), &booking); err != nil {
		if errors.Is(err, store.ErrBookingNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "booking not found"})
			return
		}
		h.logger.Printf("ERROR: updating booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update booking"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"booking": booking})

}

// DeleteBooking deletes a booking by ID
func (h *BookingsHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking id"})
		return
	}
	if err := h.bookingStore.DeleteBooking(r.Context(), uint(id)); err != nil {
		if errors.Is(err, store.ErrBookingNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "booking not found"})
			return
		}
		h.logger.Printf("ERROR: deleting booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete booking"})
		return
	}
}
