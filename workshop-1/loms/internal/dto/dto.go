package dto

import "route256/loms/internal/models"

// CreateOrderInfo is DTO for order creation
type CreateOrderInfo struct {
	Items []models.ItemOrderInfo
	models.DeliveryOrderInfo
}
