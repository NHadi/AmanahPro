package schedulers

import (
	"log"

	"github.com/robfig/cron/v3"
)

// ConfigureScheduler sets up a periodic task for reconciliation.
func ConfigureReconciliationScheduler(scheduler *cron.Cron, performReconciliation func() error) {
	_, err := scheduler.AddFunc("@every 15m", func() {
		log.Println("Starting periodic reconciliation...")
		if err := performReconciliation(); err != nil {
			log.Printf("Reconciliation failed: %v", err)
		} else {
			log.Println("Reconciliation completed successfully")
		}
	})
	if err != nil {
		log.Fatalf("Failed to configure scheduler: %v", err)
	}

	go scheduler.Start()
	log.Println("Scheduler started")
}
