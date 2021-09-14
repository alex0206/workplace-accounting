package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alex0206/workplace-accounting/internal/e"
	"github.com/alex0206/workplace-accounting/internal/model"

	"github.com/gorilla/mux"
)

// WorkplaceService describe workplace service
type WorkplaceService interface {
	Add(ctx context.Context, workplaceInfo *model.WorkplaceInfo) e.Error
	Delete(ctx context.Context, ID int) e.Error
}

// WorkplaceHandler describe workplace handler
type WorkplaceHandler struct {
	service WorkplaceService
}

// NewWorkplaceHandler create a new workplace handler
func NewWorkplaceHandler(s WorkplaceService) *WorkplaceHandler {
	return &WorkplaceHandler{service: s}
}

// Add handle of adding new workplace
func (h *WorkplaceHandler) Add(w http.ResponseWriter, r *http.Request) {
	var request model.WorkplaceInfo
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %s", err), http.StatusBadRequest)
		return
	}

	if err := h.service.Add(r.Context(), &request); err != nil {
		http.Error(w, fmt.Sprintf("error to add workplace info: %s", err.Error()), err.Code())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete handle of deleting the workplace
func (h *WorkplaceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		http.Error(w, fmt.Sprintf("error to add workplace info: %s", err.Error()), err.Code())
		return
	}

	w.WriteHeader(http.StatusOK)
}
