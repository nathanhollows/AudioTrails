package models

import "time"

// Team holds the team specific information for a given team
type Team struct {
	Clues         []Clue
	Code          string
	LastSeen      time.Time
	Solved        []int
	Unlocked      []int
	UnlockedCount int
	Events        []string
	Message       string
	Status        string
}
