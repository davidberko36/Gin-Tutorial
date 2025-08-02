package models

import "time"


type Task struct {
	ID  int64 `db:"id"`
	Title string `db:"title"`
	Description string `db:"description"`  // pending/done
	Status string `db:"status"`
	OwnerID int64 `db:"owner_id"`
	DueDate time.Time `db:"due_date"`
	CreatedAt time.Time `db:"created_at"`
}