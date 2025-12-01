package suggestions

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type CalculateTaskStatus string

const (
	CalculateTaskStatusPending   CalculateTaskStatus = "pending"
	CalculateTaskStatusRunning   CalculateTaskStatus = "running"
	CalculateTaskStatusCompleted CalculateTaskStatus = "completed"
	CalculateTaskStatusFailed    CalculateTaskStatus = "failed"
)

type CalculateTask struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Status    CalculateTaskStatus
	ErrorText string
	CreatedAt time.Time
}

type UpdateCalculateTaskFn func(ctx context.Context, task *CalculateTask) *CalculateTask
