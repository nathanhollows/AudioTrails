package game

import "errors"

// Team holds the team specific information for a given team
type Team struct {
	Code string
	Game Game
}

// Manager holds each of the teams
type Manager struct {
	Teams []Team
}

// Game holds the clues, challenges, and opportunities
type Game struct {
	Clue     int
	Ordering []int
}

// Clues is the list of clues available to all players.
var Clues []Clue

// Clue holds the next clue and any challenges associated with it
type Clue struct {
	Code      string
	Title     string
	Text      string
	Challenge string
}

// GetClue returns whether or not the clue code is valid.
func (m Manager) GetClue(code string) (Clue, error) {
	for _, clue := range Clues {
		if clue.Code == code {
			return clue, nil
		}
	}
	return Clue{}, errors.New("Clue does not exist")
}

// CheckTeam returns whether or not the team code is valid.
func (m Manager) CheckTeam(code string) bool {
	for _, clue := range Clues {
		if clue.Code == code {
			return true
		}
	}
	return false
}

// Get returns a Team given a code
func (m Manager) Get(code string) Team {
	return Team{}
}

func init() {
}
