package server

import (
	"context"
	"log/slog"
	"net/http"
	"route256/loms/internal/dto"
	"route256/loms/internal/models"
)

type LOMS interface {
	CreateOrder(ctx context.Context, userID models.UserID, info dto.CreateOrderInfo) (models.Order, error)
}

type Server struct {
	loms   LOMS
	router *http.ServeMux
}

func New(loms LOMS, router *http.ServeMux) *Server {
	slog.Info("Creating new server")
	return &Server{loms: loms, router: router}
}

func (s *Server) Run(address string) error {
	slog.Info("Starting server", slog.String("address", address))
	return http.ListenAndServe(address, s.router)
}
