package wms

import (
	"context"
	"errors"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
	oms "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/usecase/OMS"
)

type wmsService struct {
	/*
		HTTP, gRPC, XML, ... client
	*/
}

// Check that we implemet contract for usecase
var _ oms.WarehouseManagementSystem = (*wmsService)(nil)

// NewWMSService - returns WMS service adapter
func NewWMSService( /* ... */ ) *wmsService {
	return &wmsService{
		/* ... */
	}
}

func (r *wmsService) ReserveStocks(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error {
	return errors.New("unimplemented")
}
