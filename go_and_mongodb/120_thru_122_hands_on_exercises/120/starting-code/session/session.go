package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/models"

	"github.com/google/uuid"
)

type Session struct {
	inMem     inMem
	fileStore *os.File
}

type inMem map[uuid.UUID]*models.User

// New provides a new session which comes with an in-memory store
// and file system storage
func New() *Session {
	var users inMem = make(inMem)
	um, err := newFileStore(users)
	if err != nil {
		panic(err)
	}

	return &Session{users, um}
}

func newFileStore(users inMem) (*os.File, error) {
	um, err := os.Create("userMap.json")
	if err != nil {
		return nil, err
	}
	log.Println("'userMap.json' created")

	json.NewDecoder(um).Decode(&users)

	return um, um.Close()
}

func (s *Session) CreateUser(u *models.User) error {
	u.ID = uuid.New()
	fmt.Println(u)
	for {
		if _, ok := s.inMem[u.ID]; !ok {
			s.inMem[u.ID] = u
			break
		} else {
			u.ID = uuid.New()
		}
	}
	return s.writeToFile()
}

func (s *Session) GetUser(id uuid.UUID) (*models.User, error) {
	if u, ok := s.inMem[id]; ok {
		return u, nil
	}

	return nil, errors.New("user not found")
}

func (s *Session) DeleteUser(id uuid.UUID) error {
	_, ok := s.inMem[id]
	if !ok {
		return errors.New("user not found")
	}

	delete(s.inMem, id)

	return s.writeToFile()
}

func (s *Session) writeToFile() error {
	f, err := os.OpenFile("userMap.json", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		f.Close()
		return fmt.Errorf("unable able to write to file: %s", err.Error())
	}
	if err != nil {
		f.Close()
		return err
	}
	err = json.NewEncoder(f).Encode(s.inMem)
	if err != nil {
		f.Close()
		return err
	}

	return f.Close()
}
