package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	Title       string
	Description string
	Category    string
	Author      string
	Datetime    time.Time
	Location    string
	City        string
}