package services


import (
	"taskmaster/internal/models"
	"github.com/jmoiron/sqlx"
)

type taskService struct {
	db *sqlx.DB
}

func NewTaskService(db *sqlx.DB) TaskService{
	return &taskService{db: db}
}

func (s *taskService) CreateTask(t *models.Task) error {
	_, err := s.db.NamedExec(`INSERT INTO tasks (title, description, status, owner_id, due_date, created_at)
	VALUES(:title, :description, :status, :owner_id, :due_date, :created_at)`, t)
	return err
}

func (s *taskService) GetTask(id int64) (*models.Task, error) {
	var t models.Task
	err := s.db.Get(&t, "SELECT * FROM tasks WHERE id=$1", id)
	return &t, err
}

func (s *taskService) CompleteTask(id int64) error {
	_, err := s.db.Exec("UPDATE tasks SET status='done' WHERE id=$1", id)
	return err
}

func (s *taskService) CleanupExpired() (int64, error) {
	res, err := s.db.Exec("DELETE FROM tasks WHERE due_date < NOW() AND status != 'done'")
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}