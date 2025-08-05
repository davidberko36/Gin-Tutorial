package main


import (
	"log"
	"net/http"


	"github.com/go-chi/chi/v5"
	"taskmaster/internal/background"
	"taskmaster/internal/config"
	"taskmaster/internal/db"
	"taskmaster/internal/handlers"
	"taskmaster/internal/middleware"
	"taskmaster/internal/services"
)

func main() {
	cfg := config.Load()
	database, err  := db.Connect(cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	taskSvc := services.NewTaskService(database)
	taskHandler := &handlers.TaskHandler{Service :taskSvc}
	background.StartCleanup(taskSvc)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/tasks", taskHandler.Create)
	r.Get("/tasks", taskHandler.GetByID)

	log.Printf("Starting server on port %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, r)
}