package storage

import (
	"fmt"

	"github.com/buzruk/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(fileName, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetById(id uuid.UUID) (*file.File, error) {
	receivedFile, ok := s.files[id]
	if !ok {
		//return nil, errors.New("file not found")
		// return nil, errors.New(fmt.Errorf("file %v not found", id).Error())
		return nil, fmt.Errorf("file %v not found", id)
	}

	return receivedFile, nil

}
