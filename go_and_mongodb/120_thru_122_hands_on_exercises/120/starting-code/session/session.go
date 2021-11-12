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
	inMem     map[uuid.UUID]*models.User
	fileStore *os.File
}

func New() *Session {
	var users map[uuid.UUID]*models.User = make(map[uuid.UUID]*models.User)
	um, err := os.Open("userMap.json")
	if err != nil {
		log.Println(err.Error())
		log.Println("Proceeding with run, using in-memory non-nil map")
		um, err = os.Create("userMap.json")
		if err != nil {
			log.Fatal("Unable to create backup file: ", err.Error())
		}
		log.Println("'userMap.json' created")
	} else {
		json.NewDecoder(um).Decode(&users)
	}
	defer um.Close()

	return &Session{users, um}
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
