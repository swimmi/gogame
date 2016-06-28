package models

import (
	"time"
)

type Log struct {
	Type      string
	Npc       int
	Item      int
	Count     int
	Timestamp time.Time
}
