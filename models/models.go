package models

import (
	"time"

	"github.com/google/uuid"
)

type Gitevent struct {
	Uuid         uuid.UUID
	Event        string
	Eventid      string
	Branch       string
	Url          string
	Authorname   string
	Authormail   string
	DoneAt       time.Time
	Repository   string
	Addedfile    string
	Modifiedfile string
	Removedfile  string
	Message      string
}
