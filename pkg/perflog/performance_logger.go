package perflog

import (
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PerfLogger struct {
	Name   string
	ID     uuid.UUID
	Start  time.Time
	Clock  time.Time
	Logger *logrus.Logger
}

// NewPerformanceLogger will init new PerfLogger to check duration logic logging
func NewPerformanceLogger(logger *logrus.Logger, name ...string) *PerfLogger {
	pl := PerfLogger{
		ID:     uuid.New(),
		Start:  time.Now(),
		Clock:  time.Now(),
		Logger: logger,
	}
	if len(name) > 0 {
		pl.Name = name[0]
		logger.Printf("INIT NEW PERFLOG Name=%s AT %s", pl.Name, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		logger.Printf("INIT NEW PERFLOG ID=%s AT %s", pl.ID, time.Now().Format("2006-01-02 15:04:05"))
	}
	return &pl
}

// Log will write the duration of the process to logger
func (p *PerfLogger) Log(msg string) {
	now := time.Now()
	responseDuration := (float64(now.UnixNano()) - float64(p.Clock.UnixNano())) / 1000000000
	clockDuration := (float64(now.UnixNano()) - float64(p.Start.UnixNano())) / 1000000000
	p.Logger.Printf("[PERFLOG %s] Clock=%fs Dur=%fs => %s", p.Name, clockDuration, responseDuration, msg)
	p.Clock = now
}

func (p *PerfLogger) Done() {
	now := time.Now()
	clockDuration := (float64(now.UnixNano()) - float64(p.Start.UnixNano())) / 1000000000
	p.Logger.Printf("[PERFLOG %s DONE] Clock=%fs", p.Name, clockDuration)
}
