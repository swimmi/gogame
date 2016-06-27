package models

import (
	"time"
)

type PresentLog struct {
	Npc       int
	Item      int
	Count     int
	Timestamp time.Time
}
