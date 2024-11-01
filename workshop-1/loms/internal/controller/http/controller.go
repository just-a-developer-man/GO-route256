package controller_http

import "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/usecase"

type Usecases struct {
	usecase.OrderManagementSystem // OMS interface
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
