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

type CategoryHandler struct {
	categoryStore store.CategoryStore
	logger        *log.Logger
}

func NewCategoryHandler(store store.CategoryStore, logger *log.Logger) *CategoryHandler {
	return &CategoryHandler{
		categoryStore: store,
		logger:        logger,
	}
}

// CreateCategory creates a new category
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.logger.Printf("ERROR: decoding category: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	// Basic validation
	category.Name = strings.TrimSpace(category.Name)
	if category.Name == "" {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "category name is required"})
		return
	}

	if err := h.categoryStore.CreateCategory(r.Context(), &category); err != nil {
		h.logger.Printf("ERROR: creating category: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create category"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"category": category})
}

// GetCategories fetches all categories with optional sorting and pagination
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))
	orderBy := q.Get("order_by")
	if orderBy == "" {
		orderBy = "created_at asc" // default sorting
	}

	categories, err := h.categoryStore.GetCategories(r.Context(), limit, offset, orderBy)
	if err != nil {
		h.logger.Printf("ERROR: fetching categories: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve categories"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"categories": categories})
}

// GetCategoryByID fetches a single category
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid category id"})
		return
	}

	category, err := h.categoryStore.GetCategory(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrCategoryNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "category not found"})
			return
		}
		h.logger.Printf("ERROR: fetching category: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve category"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"category": category})
}

// UpdateCategory updates an existing category
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid category id"})
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.logger.Printf("ERROR: decoding category: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	// Basic validation
	category.Name = strings.TrimSpace(category.Name)
	if category.Name == "" {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "category name is required"})
		return
	}

	category.ID = uint(id)

	if err := h.categoryStore.UpdateCategory(r.Context(), &category); err != nil {
		if errors.Is(err, store.ErrCategoryNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "category not found"})
			return
		}
		h.logger.Printf("ERROR: updating category: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update category"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"category": category})
}

// DeleteCategory deletes a category by ID
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid category id"})
		return
	}

	if err := h.categoryStore.DeleteCategory(r.Context(), uint(id)); err != nil {
		if errors.Is(err, store.ErrCategoryNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "category not found"})
			return
		}
		h.logger.Printf("ERROR: deleting category: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete category"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "category deleted"})
}
