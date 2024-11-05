package models

type OrderStatus string

const (
	StatusNew             OrderStatus = "new"
	StatusAwaitingPayment OrderStatus = "awaiting payment"
	StatusFailed          OrderStatus = "failed"
	StatusPayed           OrderStatus = "payed"
	StatusCancelled       OrderStatus = "cancelled"
)

// Order - заказ
type Order struct {
	ID     OrderID         // ID заказа
	UserID UserID          // ID пользователя (чей заказ)
	Items  []ItemOrderInfo // Информация о составе заказа
}

// ItemOrderInfo - информация о составе заказа
type ItemOrderInfo struct {
	SKU      uint32 // SKU
	Quantity uint16 // количество SKU
}
