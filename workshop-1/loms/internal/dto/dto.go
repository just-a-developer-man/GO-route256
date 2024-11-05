package dto

import "github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"

// CreateOrderInfo - DTO for transferring needed data to business logic
type CreateOrderInfo struct {
	Items []models.ItemOrderInfo
}

// OrderInfo - DTO for trnasferring information about order
type OrderInfo struct {
	models.Order
	models.OrderStatus
}
