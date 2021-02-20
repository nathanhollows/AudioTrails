package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Team holds the team specific information for a given team
type Team struct {
	Code     string
	LastSeen time.Time
	Solved   []int
	Unlocked []int
	Events   []string
}

// Manager holds each of the teams
type Manager struct {
	Teams []Team
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
	for _, team := range m.Teams {
		if team.Code == code {
			return true
		}
	}
	return false
}

// GetTeam returns whether or not the team code is valid.
func (m Manager) GetTeam(code string) Team {
	for _, team := range m.Teams {
		if team.Code == code {
			return team
		}
	}
	return Team{}
}

var symbols = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func newCode() string {
	b := make([]rune, 5)
	for i := 0; i < 5; i++ {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

// CreateTeams will create however many teams are asked for.
// Num must be > 0
func (m *Manager) CreateTeams(num int) {
	for i := 0; i < num; i++ {
		m.Teams = append(m.Teams,
			Team{
				Code:     newCode(),
				Solved:   []int{},
				Unlocked: []int{rand.Intn(len(Clues)), rand.Intn(len(Clues))},
			},
		)
		fmt.Println(m.Teams[i].Code)
	}
}

// Get returns a Team given a code
func (m Manager) Get(code string) Team {
	return Team{}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	// Clues written by the wonderful Tamika!
	Clues = append(Clues, Clue{"QWOP4", "St Daves", "Named after a saint but if you have a class here you will know, saint-like is the least like what your grades will show", ""})
	Clues = append(Clues, Clue{"DFG5J", "Student Health", "Pregnancy tests and STI checks please? Book an appointment here to put your mind at ease", ""})
	Clues = append(Clues, Clue{"9LCZ4", "Campus Watch", "0800 479 5000. Save the contact.", ""})
	Clues = append(Clues, Clue{"63FJC", "Mellor Labs", "A pretty new building with chemistry inside, find the molecules to know what it hides.", ""})
	Clues = append(Clues, Clue{"9CV5A", "Radio One", "Their card gives you a bunch of sweet deals, tune in to 91fm to get in your feels", ""})
	Clues = append(Clues, Clue{"WM6SQ", "Law Library", "A library for rules and cases of times passed, climb 6-9 flights of stairs if you can be assed.", ""})
	Clues = append(Clues, Clue{"IOHRU", "Zoology Department", "Mammals, reptiles, amphibians you hear? Their kind is studied in this department I swear.", ""})
	Clues = append(Clues, Clue{"GOG79", "Black Sale House", "Our Very own Sam Leaper belongs to a band, whose name is the same as this place I understand (note: check out his poster in locals HQ)", ""})
	Clues = append(Clues, Clue{"NTODH", "Otago Business School", "Chads and Brads learn about stocks I wager, and something about hedges? (IDK im an english major)", ""})
	Clues = append(Clues, Clue{"LDQJQ", "Touchstone by the Leith", "This place is likened to a moat, and reminds me of a movie quote: 'that is a nice boulder - Donkey'", ""})
	Clues = append(Clues, Clue{"QT7KH", "Union Lawn", "A favourite place of our resident ducks, this lawn is where they give no fucks. Still stumped? Well worry no more, here you can usually find dumplings galore.", ""})
	Clues = append(Clues, Clue{"3CZ3I", "Burns Lecture Theatres", "Named after our old pal Robbie in the octagon, here is where a love of words can be built on.", ""})
	Clues = append(Clues, Clue{"B0Y2H", "Castle Lecture Theatres", "Though you may not find any knights here, this Castle is still home to things that should be feared…", ""})
	Clues = append(Clues, Clue{"I0AHP", "Botany Department", "Rhymes aside, may I just ask, can these experts revive my houseplants?", ""})
	Clues = append(Clues, Clue{"2YXNQ", "Central Library", "This library is large and perhaps the most popular, where basic bitches study, according to critic I gather.", ""})
	Clues = append(Clues, Clue{"C8P8S", "Microbiology Building", "Study something small and living too? This building is the place for you", ""})
	Clues = append(Clues, Clue{"E0MBX", "School of Physiotherapy", "Is it you neck, your back, or something else that cracks? Head to this school to get back on track.", ""})
	Clues = append(Clues, Clue{"ZEEIB", "Locals HQ", "Third floor of union where you can go, for freshers who live in flats or at home.", ""})
	Clues = append(Clues, Clue{"2SNU5", "Math / Stats Department", "I gave up on this subject long ago, but if you like numbers this place is your home.", ""})
	Clues = append(Clues, Clue{"BS0MC", "Union Grill", "If a craving for burgers, fries, and iced coffee arrives, this grill on campus has what you need to survive.", ""})
	Clues = append(Clues, Clue{"O0N66", "Science Library", "This library has diagrams and technical terms, so long to all the fiction from which we once learned. But as if that sting wasnt enough, the 7am-11pm hours will make you feel rough.", ""})
	Clues = append(Clues, Clue{"FR4SQ", "OUSA Clubs and Socs", "Their dollar-deal lunches went up in price this year. Worth it? Ill let you judge here.", ""})
	Clues = append(Clues, Clue{"H3903", "Critic Office", "Have an affinity for writing, drawing, or simply an opinion? Here, a love for journalism and magazines will make you fit in.", ""})
	Clues = append(Clues, Clue{"4IYE5", "Owheo Building", "They put the it in IT and what else I dont really know, but the computer angels frequent here, to learn at Owheo.", ""})
	Clues = append(Clues, Clue{"CFEK9", "Archway Lecture Theatres", "Consult you map for this one might be a little tricky, find the anagram of crawhAy and you should be there in a jiffy.", ""})
}
