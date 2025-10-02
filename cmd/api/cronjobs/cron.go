package cronjobs

import (
	"fmt"
	"time"
)

func Run() {
	// Initialize cron manager (enable seconds precision)
	cm := New(true)

	// Add a job every 5 seconds
	cm.AddCron("@every 5s", func() {
		fmt.Println("Task 1 running at", time.Now())
	})

	// Add a job every minute
	cm.AddCron("0 */1 * * * *", func() {
		fmt.Println("every minute cron", time.Now())
	})

	// Add a job daily at midnight
	cm.AddCron("@daily", func() {
		fmt.Println("Daily task running at", time.Now())
	})

	// Keep the program running
	select {}

}
