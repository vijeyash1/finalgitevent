package models

import (
	"time"

	"github.com/google/uuid"
)

type Gitevent struct {
	Uuid       uuid.UUID
	Event      string
	DoneBy     string
	DoneAt     time.Time
	Repository string
	Stat       string
	Message    string
}
