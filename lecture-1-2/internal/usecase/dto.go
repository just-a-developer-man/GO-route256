package usecase

import (
	"github.com/GO-route256/lecture-1-2/internal/models"
)

// CreateOrderInputInfo - DTO заказа (для создания заказа)
type CreateOrderInfo struct {
	Items                    []models.ItemOrderInfo // Товары в заказе
	models.DeliveryOrderInfo                        // Информация о доставке
}
