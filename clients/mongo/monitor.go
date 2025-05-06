package gwm_mongo

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/event"
)

// MongoCommandMonitor implements command monitoring
type MongoCommandMonitor struct {
	operations sync.Map
	threshold  time.Duration
}

// NewMongoCommandMonitor creates a new command monitor that implements event.CommandMonitor
func NewMongoCommandMonitor() *event.CommandMonitor {
	m := &MongoCommandMonitor{
		threshold:  time.Duration(slowQueryThreshold) * time.Millisecond,
		operations: sync.Map{},
	}

	// Return a properly configured event.CommandMonitor
	return &event.CommandMonitor{
		Started:   m.Started,
		Succeeded: m.Succeeded,
		Failed:    m.Failed,
	}
}

// Started is called when a command starts execution
func (m *MongoCommandMonitor) Started(ctx context.Context, evt *event.CommandStartedEvent) {
	m.operations.Store(evt.RequestID, time.Now())
}

// Succeeded is called when a command completes successfully
func (m *MongoCommandMonitor) Succeeded(ctx context.Context, evt *event.CommandSucceededEvent) {
	slog.DebugContext(ctx, "MongoDB command succeeded",
		"command", evt.CommandName,
		"duration", evt.Duration,
		"database", evt.DatabaseName,
		"connectionID", evt.ConnectionID,
		"requestID", evt.RequestID,
	)
	m.checkDuration(ctx, evt.RequestID, evt.CommandName, evt.DatabaseName, evt.ConnectionID)
}

// Failed is called when a command execution fails
func (m *MongoCommandMonitor) Failed(ctx context.Context, evt *event.CommandFailedEvent) {
	slog.WarnContext(ctx, "MongoDB command failed",
		"command", evt.CommandName,
		"error", evt.Failure,
		"database", evt.DatabaseName,
		"connectionID", evt.ConnectionID,
		"requestID", evt.RequestID,
	)
	m.checkDuration(ctx, evt.RequestID, evt.CommandName, evt.DatabaseName, evt.ConnectionID)
}

// checkDuration checks if the operation exceeded the threshold and logs it
func (m *MongoCommandMonitor) checkDuration(
	ctx context.Context,
	requestID int64, commandName, dbName string, connID string,
) {
	startTime, exists := m.operations.Load(requestID)
	if !exists {
		return
	}

	duration := time.Since(startTime.(time.Time))
	m.operations.Delete(requestID)

	if duration >= m.threshold {
		slog.WarnContext(ctx, "MongoDB slow query detected",
			"command", commandName,
			"duration_ms", duration.Milliseconds(),
			"database", dbName,
			"connectionID", connID,
			"threshold_ms", m.threshold.Milliseconds(),
		)
	}
}
