package controller_http

import "github.com/GO-route256/classroom-8/students/lecture-1-2/internal/usecase"

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
