package usecase

import (
	"context"

	"github.com/just-a-developer-man/GO-route256/lecture-1-2/internal/models"
)

// interfaces.go: Декларируем бизнес функциональность

type OrderManagementSystem interface {
	CreateOrder(ctx context.Context, userID models.UserID, info CreateOrderInfo) (models.Order, error)
	/* ... */
}
