package notes

import (
	"context"
	"sync"

	"github.com/just-a-developer-man/GO-route256/workshop-2/internal/model"
)

type Repository struct {
	notes []*model.Note
	mx    sync.Mutex
}

func NewRepository() *Repository {
	return &Repository{}
}

func (s *Repository) SaveNote(ctx context.Context, note *model.Note) (int, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	note.Id = len(s.notes)
	s.notes = append(s.notes, note)
	return len(s.notes) - 1, nil
}

func (s *Repository) ListNotes(ctx context.Context) ([]*model.Note, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.notes, nil
}
