package models

import "time"

// Team holds the team specific information for a given team
type Team struct {
	Code          string
	LastSeen      time.Time
	UnlockedCount int
	Message       string
	Status        string
}
