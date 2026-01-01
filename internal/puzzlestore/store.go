package puzzlestore

import (
	"sync"

	"github.com/google/uuid"
)

//this is temp
// will be replaced with actual mongo/sql instance later

type Store struct {
	lock sync.RWMutex
	store map[uuid.UUID]*Puzzle
}

type Puzzle struct {
	ID uuid.UUID
	GameState []*Position;
	PlayerSide string
	Moves []string
	Level int32
}

type Position struct {
	Piece string
	Placement string
}

func (s *Store) CreateNewPuzzle (newPuzzle *Puzzle) uuid.UUID {
	defer s.lock.Unlock()
	newUUID := uuid.New()
	s.lock.Lock()
	s.store[newUUID] = newPuzzle;
	return newUUID
}

func (s *Store) GetPuzzleByID (id uuid.UUID) *Puzzle {
	defer s.lock.RUnlock()
	s.lock.RLock()
	if _, ok := s.store[id]; ok {
		return s.store[id]
	}
	return nil
}

func (s *Store) DeletePuzzleByID (id uuid.UUID) {
	defer s.lock.Unlock()
	s.lock.Lock()
	delete(((*s).store), id)
}

func (s *Store) GetPuzzles (level int32) []*Puzzle {
	var rN []*Puzzle = make([]*Puzzle, 0)
	for _, v := range s.store {
		if v.Level == level {
			rN = append(rN, v)
		}
		if len(rN) > 4 {
			return rN
		}
	}
	return rN
}

func NewStore () *Store {
	return &Store{}
}