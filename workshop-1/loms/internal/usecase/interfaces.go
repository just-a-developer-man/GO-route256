package usecase

import (
	"context"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
)

// interfaces.go: Декларируем бизнес функциональность

type OrderManagementSystem interface {
	CreateOrder(ctx context.Context, userID models.UserID, info CreateOrderInfo) (models.OrderID, error)
	OrderByID(ctx context.Context, orderID models.OrderID) (models.Order, error)
	/* ... */
}
