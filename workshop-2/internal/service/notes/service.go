package notes

import (
	"context"

	"github.com/just-a-developer-man/GO-route256/workshop-2/internal/model"
	"github.com/just-a-developer-man/GO-route256/workshop-2/internal/repository/notes"
)

type Service struct {
	repo *notes.Repository
}

func NewService() *Service {
	return &Service{notes.NewRepository()}
}

func (s *Service) SaveNote(ctx context.Context, note *model.Note) (int, error) {
	return s.repo.SaveNote(ctx, note)
}

func (s *Service) ListNotes(ctx context.Context) ([]*model.Note, error) {
	return s.repo.ListNotes(ctx)
}
