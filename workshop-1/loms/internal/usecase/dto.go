package usecase

import (
	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/models"
)

// CreateOrderInputInfo - DTO заказа (для создания заказа)
type CreateOrderInfo struct {
	Items                    []models.ItemOrderInfo // Товары в заказе
	models.DeliveryOrderInfo                        // Информация о доставке
}
