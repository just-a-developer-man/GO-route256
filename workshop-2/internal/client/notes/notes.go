package notes

import (
	"context"

	"github.com/just-a-developer-man/GO-route256/workshop-2/internal/converter/client"
	"github.com/just-a-developer-man/GO-route256/workshop-2/internal/model"
	notes_v1 "github.com/just-a-developer-man/GO-route256/workshop-2/pkg/api/notes/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	notes_v1.NotesClient
}

func NewClient(c notes_v1.NotesClient) *Client {
	return &Client{NotesClient: c}
}

func (s *Client) SaveNote(ctx context.Context, note *model.Note) (int, error) {
	rs, err := s.NotesClient.SaveNote(ctx, client.NoteToReq(note))
	if err != nil {
		return 0, err
	}
	return int(rs.GetNoteId()), nil
}

func (s *Client) ListNotes(ctx context.Context) ([]*model.Note, error) {
	rs, err := s.NotesClient.ListNotes(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return client.NotesFromResp(rs), nil
}
