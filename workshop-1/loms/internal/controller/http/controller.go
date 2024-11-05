package controller_http

import (
	"context"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/dto"
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
)

type OrderManagementSystem interface {
	CreateOrder(ctx context.Context, userID models.UserID, info dto.CreateOrderInfo) (models.OrderID, error)
	OrderByID(ctx context.Context, orderID models.OrderID) (dto.OrderInfo, error)
	/* ... */
}

type Usecases struct {
	OrderManagementSystem // OMS interface
}

// Controller - is controller/delivery layer
type Controller struct {
	Usecases
	/* ... */
}

// NewController - returns Controller
func NewController(us Usecases) *Controller {
	return &Controller{
		Usecases: us,
	}
}
