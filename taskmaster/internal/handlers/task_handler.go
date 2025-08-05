package handlers


import (
	"encoding/json"
	"net/http"
	"strconv"
	"taskmaster/internal/models"
	"taskmaster/internal/services"
	"time"
)

type TaskHandler struct {
	Service services.TaskService
}


func (h *TaskHandler) Create(w http.ResponseWriter,r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	t.Status = "pending"
	t.CreatedAt = time.Now()
	if err := h.Service.CreateTask(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}


func (h *TaskHandler) getByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	task, err := h.Service.GetTask(id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}