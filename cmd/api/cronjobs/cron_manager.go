package cronjobs

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

// CronManager holds the cron scheduler and ensures thread-safety
type CronManager struct {
	cron *cron.Cron
	mu   sync.Mutex
}

// New creates a new CronManager
func New(enableSeconds bool) *CronManager {
	var c *cron.Cron
	if enableSeconds {
		c = cron.New(cron.WithSeconds())
	} else {
		c = cron.New()
	}

	manager := &CronManager{
		cron: c,
	}

	manager.cron.Start() // start scheduler immediately
	return manager
}

// AddCron adds a cron job with a schedule and function
func (m *CronManager) AddCron(schedule string, job func()) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, err := m.cron.AddFunc(schedule, job)
	if err != nil {
		return fmt.Errorf("failed to add cron job: %w", err)
	}
	return nil
}

// Stop stops the cron scheduler
func (m *CronManager) Stop() {
	m.cron.Stop()
}
