package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, data []byte) (*File, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	file := File{
		ID:   id,
		Name: name,
		Data: data,
	}

	return &file, nil
}

func (f *File) String() string {
	return fmt.Sprintf("ID: %v, Name: %v, Data: \"%v\"", f.ID, f.Name, string(f.Data))
}
