package background


import (
	"fmt"
	"time"
	"taskmaster/internal/services"
)


func StartCleanup(svc services.TaskService){
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker .C {
			count, err := svc.CleanupExpired()
			if err != nil{
				fmt.Println("Clean up failed:", err)
				continue
			}
			fmt.Printf("Cleaned up %d tasks\n", count)
		}
	}()
}