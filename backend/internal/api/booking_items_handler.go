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

type BookingItemsHandler struct {
	bookingItemStore store.BookingItemStore
	logger           *log.Logger
}

func NewBookingItemsHandler(store store.BookingItemStore, logger *log.Logger) *BookingItemsHandler {
	return &BookingItemsHandler{
		bookingItemStore: store,
		logger:           logger,
	}
}

// CreateBookingItem creates a new booking item
func (h *BookingItemsHandler) CreateBookingItem(w http.ResponseWriter, r *http.Request) {
	var bookingItem models.BookingItem
	if err := json.NewDecoder(r.Body).Decode(&bookingItem); err != nil {
		h.logger.Printf("ERROR: decoding booking item: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	if err := h.bookingItemStore.CreateBookingItem(r.Context(), &bookingItem); err != nil {
		h.logger.Printf("ERROR: creating booking item: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create booking item"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"booking item": bookingItem})
}

// GetBookingItems fetches all booking items with optional sorting and pagination
func (h *BookingItemsHandler) GetBookingItems(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))
	orderBy := q.Get("order_by")
	if orderBy == "" {
		orderBy = "created_at asc" // default sorting
	}

	bookingItems, err := h.bookingItemStore.GetBookingItems(r.Context(), limit, offset, orderBy)
	if err != nil {
		h.logger.Printf("ERROR: fetching booking items: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve booking items"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"booking items": bookingItems})

}

// GetBookingItemByID fetches a single booking item
func (h *BookingItemsHandler) GetBookingItemByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking item id"})
		return
	}

	bookingItem, err := h.bookingItemStore.GetBookingItem(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrBookingItemNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "booking item not found"})
			return
		}
		h.logger.Printf("ERROR: fetching booking item: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve booking item"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"booking item	": bookingItem})
}

// UpdateBookingItem updates an existing booking item
func (h *BookingItemsHandler) UpdateBookingItem(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking item id"})
		return
	}

	var bookingItem models.BookingItem
	if err := json.NewDecoder(r.Body).Decode(&bookingItem); err != nil {
		h.logger.Printf("ERROR: decoding booking item: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	bookingItem.ID = uint(id)
	if err := h.bookingItemStore.UpdateBookingItem(r.Context(), &bookingItem); err != nil {
		if errors.Is(err, store.ErrBookingItemNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "booking item not found"})
			return
		}
		h.logger.Printf("ERROR: updating booking item: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update booking item"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"booking item": bookingItem})

}

// DeleteBookingItem deletes a booking item by ID
func (h *BookingItemsHandler) DeleteBookingItem(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking item id"})
		return
	}
	if err := h.bookingItemStore.DeleteBookingItem(r.Context(), uint(id)); err != nil {
		if errors.Is(err, store.ErrBookingItemNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "booking item not found"})
			return
		}
		h.logger.Printf("ERROR: deleting booking item: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete booking item"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "booking item deleted"})

}
