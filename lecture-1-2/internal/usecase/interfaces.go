package usecase

import (
	"context"

	"github.com/GO-route256/classroom-8/students/lecture-1-2/internal/models"
)

// interfaces.go: Декларируем бизнес функциональность

type OrderManagementSystem interface {
	CreateOrder(ctx context.Context, userID models.UserID, info CreateOrderInfo) (models.Order, error)
	/* ... */
}
